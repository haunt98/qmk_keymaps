package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

var (
	flagDebug     bool
	flagQMKInfo   string
	flagQMKKeymap string
	flagOutput    string
)

func init() {
	flag.BoolVar(&flagDebug, "debug", false, "Debug mode, print eveyrthing")
	flag.StringVar(&flagQMKInfo, "qmk-info", "", "QMK info json file")
	flag.StringVar(&flagQMKKeymap, "qmk-keymap", "", "QMK keymap json file")
	flag.StringVar(&flagOutput, "out", "", "Output file")
}

func main() {
	flag.Parse()

	if flagQMKInfo == "" {
		log.Fatalln("Flag --qmk-info is empty")
	}

	if flagQMKKeymap == "" {
		log.Fatalln("Flag --qmk-keymap is empty")
	}

	if flagDebug {
		log.Printf("flagQMKInfo: [%s]\n", flagQMKInfo)
		log.Printf("flagQMKKeymap: [%s]\n", flagQMKKeymap)
		log.Printf("flagOutput: [%s]\n", flagOutput)
	}

	bytesInfo, err := os.ReadFile(flagQMKInfo)
	if err != nil {
		log.Fatalln("Failed to read file", flagQMKInfo, err)
	}

	qmkInfo := QMKInfo{}
	if err := json.Unmarshal(bytesInfo, &qmkInfo); err != nil {
		log.Fatalln("Failed to json unmarshal", flagQMKInfo, err)
	}

	bytesKeymap, err := os.ReadFile(flagQMKKeymap)
	if err != nil {
		log.Fatalln("Failed to read file", flagQMKKeymap, err)
	}

	qmkKeymap := QMKKeymap{}
	if err := json.Unmarshal(bytesKeymap, &qmkKeymap); err != nil {
		log.Fatalln("Failed to json unmarshal", flagQMKKeymap, err)
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
