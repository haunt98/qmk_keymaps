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
	ErrEmptyQMKInfo = errors.New("empty QMK info")

	flagDebug         bool
	flagQMKKeyboard   string
	flagQMKInfoFle    string
	flagQMKKeymapFile string
	flagOutput        string
)

func init() {
	flag.BoolVar(&flagDebug, "debug", false, "Debug mode, print eveyrthing")
	flag.StringVar(&flagQMKKeyboard, "qmk-keyboard", "", "QMK keyboard name")
	flag.StringVar(&flagQMKInfoFle, "qmk-info-file", "", "QMK info json file")
	flag.StringVar(&flagQMKKeymapFile, "qmk-keymap-file", "", "QMK keymap json file")
	flag.StringVar(&flagOutput, "out", "", "Output file")
}

func main() {
	flag.Parse()
	if flagDebug {
		log.Printf("flagQMKInfo: [%s]\n", flagQMKInfoFle)
		log.Printf("flagQMKKeymap: [%s]\n", flagQMKKeymapFile)
		log.Printf("flagOutput: [%s]\n", flagOutput)
	}

	qmkInfo := QMKInfo{}
	// -qmk-keyboard first, then fallback to -qmk-info-file
	if flagQMKKeyboard != "" {
		var err error
		qmkInfo, err = getQMKInfo(flagQMKKeyboard, flagDebug)
		if err != nil {
			log.Fatalf("Failed to get QMK info [%s]: %s\n", flagQMKKeyboard, err)
		}
	} else {
		if flagQMKInfoFle == "" {
			log.Fatalln("Flag -qmk-info-file is empty")
		}

		bytesInfo, err := os.ReadFile(flagQMKInfoFle)
		if err != nil {
			log.Fatalf("Failed to read file [%s]: %s\n", flagQMKInfoFle, err)
		}

		if err := json.Unmarshal(bytesInfo, &qmkInfo); err != nil {
			log.Fatalf("Failed to json unmarshal [%s]: %s\n", flagQMKInfoFle, err)
		}
	}

	if flagQMKKeymapFile == "" {
		log.Fatalln("Flag -qmk-keymap is empty")
	}

	bytesKeymap, err := os.ReadFile(flagQMKKeymapFile)
	if err != nil {
		log.Fatalln("Failed to read file", flagQMKKeymapFile, err)
	}

	qmkKeymap := QMKKeymap{}
	if err := json.Unmarshal(bytesKeymap, &qmkKeymap); err != nil {
		log.Fatalln("Failed to json unmarshal", flagQMKKeymapFile, err)
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

func getQMKInfo(qmkKeyboard string, flagDebug bool) (QMKInfo, error) {
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
			if flagDebug {
				log.Printf("Failed to http get [%s]: %s\n", url, err)
			}

			continue
		}

		data, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			if flagDebug {
				log.Printf("Failed to read all body [%s]: %s\n", url, err)
			}

			continue
		}

		qmkInfo := QMKInfo{}
		if err := json.Unmarshal(data, &qmkInfo); err != nil {
			if flagDebug {
				log.Printf("Failed to json unmarshal [%s]: %s\n", url, err)
			}

			continue
		}

		if len(qmkInfo.Layouts) == 0 {
			continue
		}

		if flagDebug {
			log.Printf("Found QMK info for [%s]\n", kb)
		}

		return qmkInfo, nil
	}

	return QMKInfo{}, ErrEmptyQMKInfo
}
