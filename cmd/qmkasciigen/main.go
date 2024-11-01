package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/spf13/cast"
)

const (
	// QMK template
	qmkKeymapURL = "https://raw.githubusercontent.com/qmk/qmk_firmware/master/keyboards/%s/keymaps/%s/keymap.json"
)

// QMK template
var qmkKeyboardInfoURLs = []string{
	"https://raw.githubusercontent.com/qmk/qmk_firmware/master/keyboards/%s/keyboard.json",
	"https://raw.githubusercontent.com/qmk/qmk_firmware/master/keyboards/%s/info.json",
}

var (
	ErrEmptyQMKInfo   = errors.New("empty QMK info")
	ErrEmptyQMKKeymap = errors.New("empty QMK keymap")

	flagDebug         bool
	flagQMKKeyboard   string
	flagQMKKeymap     string
	flagQMKInfoFile   string
	flagQMKKeymapFile string
	flagOutput        string
	flagPrintOutput   bool
	flagPrintLayout   bool
	flagPrintLayer    bool
	flagAllowLayers   string
)

func init() {
	flag.BoolVar(&flagDebug, "debug", false, "Debug mode, print eveyrthing")
	flag.StringVar(&flagQMKKeyboard, "qmk-keyboard", "", "QMK keyboard name")
	flag.StringVar(&flagQMKKeymap, "qmk-keymap", "default", "QMK keymap name")
	flag.StringVar(&flagQMKInfoFile, "qmk-info-file", "", "QMK info json file")
	flag.StringVar(&flagQMKKeymapFile, "qmk-keymap-file", "", "QMK keymap json file")
	flag.StringVar(&flagOutput, "out", "", "Output file")
	flag.BoolVar(&flagPrintOutput, "print-out", false, "Print output")
	flag.BoolVar(&flagPrintLayout, "print-layout", false, "Print layout name")
	flag.BoolVar(&flagPrintLayer, "print-layer", false, "Print layer name")
	flag.StringVar(&flagAllowLayers, "allow-layers", "", "Allow layers to print aka 0 or 0,2,3. Use , to split. Empty mean all")
}

func main() {
	flag.Parse()
	if flagDebug {
		log.Printf("flagQMKKeyboard: [%s]\n", flagQMKKeyboard)
		log.Printf("flagQMKKeymap: [%s]\n", flagQMKKeymap)
		log.Printf("flagQMKInfoFile: [%s]\n", flagQMKInfoFile)
		log.Printf("flagQMKKeymapFile: [%s]\n", flagQMKKeymapFile)
		log.Printf("flagOutput: [%s]\n", flagOutput)
		log.Printf("flagPrintOutput: [%v]\n", flagPrintOutput)
		log.Printf("flagPrintLayout: [%v]\n", flagPrintLayout)
		log.Printf("flagPrintLayer: [%v]\n", flagPrintLayer)
		log.Printf("flagAllowLayers: [%s]\n", flagAllowLayers)
	}

	// 0,2,3 -> [0, 2, 3]
	allowLayers := make(map[int]struct{})
	for _, l := range strings.Split(flagAllowLayers, ",") {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		lInt := cast.ToInt(l)
		allowLayers[lInt] = struct{}{}
	}

	qmkInfo, err := wrapGetQMKInfo(flagQMKKeyboard, flagQMKInfoFile, flagDebug)
	if err != nil {
		log.Fatalln("Failed to get QMK info", err)
	}

	qmkKeymap, err := wrapGetQMKKeymap(flagQMKKeyboard, flagQMKKeymap, flagQMKKeymapFile, flagDebug)
	if err != nil {
		log.Fatalln("Failed to get QMK keymap", err)
	}

	result := Draw(
		qmkInfo.Layouts,
		qmkKeymap,
		DrawConfig{
			AllowLayers: allowLayers,
			PrintLayout: flagPrintLayout,
			PrintLayer:  flagPrintLayer,
		},
	)
	if flagPrintOutput {
		fmt.Println(result)
	}

	// Optional
	if flagOutput != "" {
		if err := os.WriteFile(flagOutput, []byte(result), 0o600); err != nil {
			log.Fatalln("Failed to write file", flagOutput, err)
		}
	}
}

// Local first, remote later
func wrapGetQMKInfo(qmkKeyboardStr, qmkInfoFile string, debug bool) (QMKInfo, error) {
	// Try to get local file
	if qmkInfoFile != "" {
		bytesInfo, err := os.ReadFile(qmkInfoFile)
		if err != nil {
			return QMKInfo{}, fmt.Errorf("os: failed to read file [%s]: %w", qmkInfoFile, err)
		}

		qmkInfo := QMKInfo{}
		if err := sonic.Unmarshal(bytesInfo, &qmkInfo); err != nil {
			return QMKInfo{}, fmt.Errorf("sonic: failed to unmarshal [%s]: %w", qmkInfoFile, err)
		}

		return qmkInfo, nil
	}

	// Try to get remote file
	if qmkKeyboardStr != "" {
		qmkInfo, err := getQMKInfo(qmkKeyboardStr, debug)
		if err != nil {
			return QMKInfo{}, err
		}

		return qmkInfo, nil
	}

	return QMKInfo{}, ErrEmptyQMKInfo
}

func getQMKInfo(qmkKeyboardStr string, debug bool) (QMKInfo, error) {
	qmkKeyboardStr = strings.TrimSpace(qmkKeyboardStr)
	if qmkKeyboardStr == "" {
		return QMKInfo{}, ErrEmptyQMKInfo
	}

	kbParts := strings.Split(qmkKeyboardStr, "/")
	for i := len(kbParts); i >= 1; i-- {
		kb := strings.Join(kbParts[:i], "/")

		for _, uTemplate := range qmkKeyboardInfoURLs {
			u := fmt.Sprintf(uTemplate, kb)

			qmkInfo, found := getQMKInfoByURL(u, debug)
			if found {
				return qmkInfo, nil
			}
		}
	}

	return QMKInfo{}, ErrEmptyQMKInfo
}

// Ignore error because we try to find where that url is
func getQMKInfoByURL(u string, debug bool) (QMKInfo, bool) {
	// nolint:noctx,gosec
	httpRsp, err := http.Get(u)
	if err != nil {
		if debug {
			log.Printf("Failed to http get [%s]: %s\n", u, err)
		}

		return QMKInfo{}, false
	}
	defer httpRsp.Body.Close()

	data, err := io.ReadAll(httpRsp.Body)
	if err != nil {
		if debug {
			log.Printf("Failed to read all body [%s]: %s\n", u, err)
		}

		return QMKInfo{}, false
	}

	qmkInfo := QMKInfo{}
	if err := sonic.Unmarshal(data, &qmkInfo); err != nil {
		if debug {
			log.Printf("Failed to sonic unmarshal [%s]: %s\n", u, err)
		}

		return QMKInfo{}, false
	}

	if len(qmkInfo.Layouts) == 0 {
		if debug {
			log.Printf("Empty layouts [%s]\n", u)
		}

		return QMKInfo{}, false
	}

	if debug {
		log.Printf("Found QMK info [%s]\n", u)
	}

	return qmkInfo, true
}

// Local first, remote later
func wrapGetQMKKeymap(qmkKeyboardStr, qmkKeymapStr, qmkKeymapFile string, debug bool) (QMKKeymap, error) {
	// Try to get local file
	if qmkKeymapFile != "" {
		bytesKeymap, err := os.ReadFile(qmkKeymapFile)
		if err != nil {
			return QMKKeymap{}, fmt.Errorf("os: failed to read file [%s]: %w", qmkKeymapFile, err)
		}

		qmkKeymap := QMKKeymap{}
		if err := sonic.Unmarshal(bytesKeymap, &qmkKeymap); err != nil {
			return QMKKeymap{}, fmt.Errorf("sonic: failed to unmarshal [%s]: %w", qmkKeymapFile, err)
		}

		return qmkKeymap, nil
	}

	// Try to get remote file
	if qmkKeyboardStr != "" && qmkKeymapStr != "" {
		qmkKeymap, err := getQMKKeymap(qmkKeyboardStr, qmkKeymapStr, debug)
		if err != nil {
			return QMKKeymap{}, err
		}

		return qmkKeymap, nil
	}

	return QMKKeymap{}, ErrEmptyQMKKeymap
}

func getQMKKeymap(qmkKeyboardStr, qmkKeymapStr string, debug bool) (QMKKeymap, error) {
	if qmkKeyboardStr == "" || qmkKeymapStr == "" {
		return QMKKeymap{}, ErrEmptyQMKKeymap
	}

	// See qmk url
	kbParts := strings.Split(qmkKeyboardStr, "/")
	for i := len(kbParts); i >= 1; i-- {
		kb := strings.Join(kbParts[:i], "/")
		u := fmt.Sprintf(qmkKeymapURL, kb, qmkKeymapStr)

		qmkKeymap, found := getQMKKeymapByURL(u, debug)
		if found {
			return qmkKeymap, nil
		}
	}

	return QMKKeymap{}, ErrEmptyQMKKeymap
}

// Ignore error because we try to find where that url is
func getQMKKeymapByURL(u string, debug bool) (QMKKeymap, bool) {
	// nolint:noctx,gosec
	httpRsp, err := http.Get(u)
	if err != nil {
		if debug {
			log.Printf("Failed to http get [%s]: %s\n", u, err)
		}

		return QMKKeymap{}, false
	}
	defer httpRsp.Body.Close()

	data, err := io.ReadAll(httpRsp.Body)
	if err != nil {
		if debug {
			log.Printf("Failed to read all body [%s]: %s\n", u, err)
		}

		return QMKKeymap{}, false
	}

	qmkKeymap := QMKKeymap{}
	if err := sonic.Unmarshal(data, &qmkKeymap); err != nil {
		if debug {
			log.Printf("Failed to sonic unmarshal [%s]: %s\n", u, err)
		}

		return QMKKeymap{}, false
	}

	if len(qmkKeymap.Layers) == 0 {
		return QMKKeymap{}, false
	}

	if debug {
		log.Printf("Found QMK keymap [%s]\n", u)
	}

	return qmkKeymap, true
}
