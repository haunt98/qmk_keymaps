package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	qmkKeyboardInfoURL = "https://raw.githubusercontent.com/qmk/qmk_firmware/master/keyboards/%s/info.json"
	qmkKeymapURL       = "https://raw.githubusercontent.com/qmk/qmk_firmware/master/keyboards/%s/keymaps/%s/keymap.json"
)

var (
	ErrEmptyQMKInfo       = errors.New("empty QMK info")
	ErrEmptyQMKKeymap     = errors.New("empty QMK keymap")
	ErrEmptyQMKInfoFile   = errors.New("empty QMK info file")
	ErrEmptyQMKKeymapFile = errors.New("empty QMK keymap file")

	flagDebug         bool
	flagQMKKeyboard   string
	flagQMKKeymap     string
	flagQMKInfoFile   string
	flagQMKKeymapFile string
	flagOutput        string
)

func init() {
	flag.BoolVar(&flagDebug, "debug", false, "Debug mode, print eveyrthing")
	flag.StringVar(&flagQMKKeyboard, "qmk-keyboard", "", "QMK keyboard name")
	flag.StringVar(&flagQMKKeymap, "qmk-keymap", "default", "QMK keymap name")
	flag.StringVar(&flagQMKInfoFile, "qmk-info-file", "", "QMK info json file")
	flag.StringVar(&flagQMKKeymapFile, "qmk-keymap-file", "", "QMK keymap json file")
	flag.StringVar(&flagOutput, "out", "", "Output file")
}

func main() {
	flag.Parse()
	if flagDebug {
		log.Printf("flagQMKKeyboard: [%s]\n", flagQMKKeyboard)
		log.Printf("flagQMKKeymap: [%s]\n", flagQMKKeymap)
		log.Printf("flagQMKInfoFile: [%s]\n", flagQMKInfoFile)
		log.Printf("flagQMKKeymapFile: [%s]\n", flagQMKKeymapFile)
		log.Printf("flagOutput: [%s]\n", flagOutput)
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
	)
	if flagDebug {
		log.Printf("Result:\n%s\n", result)
	}

	// Optional
	if flagOutput != "" {
		if err := os.WriteFile(flagOutput, []byte(result), 0o600); err != nil {
			log.Fatalln("Failed to write file", flagOutput, err)
		}
	}
}

// Remote first, fallback local later
func wrapGetQMKInfo(qmkKeyboardStr, qmkInfoFile string, debug bool) (QMKInfo, error) {
	if qmkKeyboardStr != "" {
		qmkInfo, err := getQMKInfo(qmkKeyboardStr, debug)
		if err != nil {
			if debug {
				log.Printf("Failed to get QMK info [%s]: %s\n", qmkKeyboardStr, err)
			}
		} else {
			return qmkInfo, nil
		}
	}

	if qmkInfoFile == "" {
		return QMKInfo{}, ErrEmptyQMKInfoFile
	}

	bytesInfo, err := os.ReadFile(qmkInfoFile)
	if err != nil {
		return QMKInfo{}, fmt.Errorf("os: failed to read file [%s]: %w", qmkInfoFile, err)
	}

	qmkInfo := QMKInfo{}
	if err := json.Unmarshal(bytesInfo, &qmkInfo); err != nil {
		return QMKInfo{}, fmt.Errorf("json: failed to unmarshal [%s]: %w", qmkInfoFile, err)
	}

	return qmkInfo, nil
}

func getQMKInfo(qmkKeyboard string, debug bool) (QMKInfo, error) {
	qmkKeyboard = strings.TrimSpace(qmkKeyboard)
	if qmkKeyboard == "" {
		return QMKInfo{}, ErrEmptyQMKInfo
	}

	kbParts := strings.Split(qmkKeyboard, "/")
	for i := 1; i <= len(kbParts); i++ {
		kb := strings.Join(kbParts[:i], "/")

		url := fmt.Sprintf(qmkKeyboardInfoURL, kb)

		// nolint:noctx,gosec
		httpRsp, err := http.Get(url)
		if err != nil {
			if debug {
				log.Printf("Failed to http get [%s]: %s\n", url, err)
			}

			continue
		}

		data, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			if debug {
				log.Printf("Failed to read all body [%s]: %s\n", url, err)
			}

			continue
		}

		qmkInfo := QMKInfo{}
		if err := json.Unmarshal(data, &qmkInfo); err != nil {
			if debug {
				log.Printf("Failed to json unmarshal [%s]: %s\n", url, err)
			}

			continue
		}

		if len(qmkInfo.Layouts) == 0 {
			continue
		}

		if debug {
			log.Printf("Found QMK info [%s]\n", url)
		}

		return qmkInfo, nil
	}

	return QMKInfo{}, ErrEmptyQMKInfo
}

// Remote first, fallback local later
func wrapGetQMKKeymap(qmkKeyboardStr, qmkKeymapStr, qmkKeymapFile string, debug bool) (QMKKeymap, error) {
	if qmkKeyboardStr != "" && qmkKeymapStr != "" {
		qmkKeymap, err := getQMKKeymap(qmkKeyboardStr, qmkKeymapStr, debug)
		if err != nil {
			if debug {
				log.Printf("Failed to get QMK keymap [%s] [%s]: %s\n", qmkKeyboardStr, qmkKeymapStr, err)
			}
		} else {
			return qmkKeymap, nil
		}
	}

	if qmkKeymapFile == "" {
		return QMKKeymap{}, ErrEmptyQMKKeymapFile
	}

	bytesKeymap, err := os.ReadFile(qmkKeymapFile)
	if err != nil {
		return QMKKeymap{}, fmt.Errorf("os: failed to read file [%s]: %w", qmkKeymapFile, err)
	}

	qmkKeymap := QMKKeymap{}
	if err := json.Unmarshal(bytesKeymap, &qmkKeymap); err != nil {
		return QMKKeymap{}, fmt.Errorf("json: failed to unmarshal [%s]: %w", qmkKeymapFile, err)
	}

	return qmkKeymap, nil
}

func getQMKKeymap(qmkKeyboard, qmkKeymap string, debug bool) (QMKKeymap, error) {
	qmkKeyboard = strings.TrimSpace(qmkKeyboard)
	qmkKeymap = strings.TrimSpace(qmkKeymap)
	if qmkKeyboard == "" || qmkKeymap == "" {
		return QMKKeymap{}, ErrEmptyQMKKeymap
	}

	kbParts := strings.Split(qmkKeyboard, "/")
	for i := 1; i <= len(kbParts); i++ {
		kb := strings.Join(kbParts[:i], "/")

		url := fmt.Sprintf(qmkKeymapURL, kb, qmkKeymap)

		// nolint:noctx,gosec
		httpRsp, err := http.Get(url)
		if err != nil {
			if debug {
				log.Printf("Failed to http get [%s]: %s\n", url, err)
			}

			continue
		}

		data, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			if debug {
				log.Printf("Failed to read all body [%s]: %s\n", url, err)
			}

			continue
		}

		qmkKeymap := QMKKeymap{}
		if err := json.Unmarshal(data, &qmkKeymap); err != nil {
			if debug {
				log.Printf("Failed to json unmarshal [%s]: %s\n", url, err)
			}

			continue
		}

		if len(qmkKeymap.Layers) == 0 {
			continue
		}

		if debug {
			log.Printf("Found QMK keymap [%s]\n", url)
		}

		return qmkKeymap, nil
	}

	return QMKKeymap{}, ErrEmptyQMKKeymap
}
