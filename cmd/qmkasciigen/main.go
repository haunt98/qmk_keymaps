package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func main() {
	debug := strings.EqualFold(os.Getenv("DEBUG"), "true")

	filenameInfo := os.Getenv("QMK_INFO")
	if filenameInfo == "" {
		log.Fatalln("ENV QMK_INFO empty")
	}
	if debug {
		log.Printf("filenameInfo: [%s]\n", filenameInfo)
	}

	filenameKeymap := os.Getenv("QMK_KEYMAP")
	if filenameKeymap == "" {
		log.Fatalln("ENV QMK_KEYMAP empty")
	}
	if debug {
		log.Printf("filenameKeymap: [%s]\n", filenameKeymap)
	}

	filenameOut := os.Getenv("OUT")
	if debug {
		log.Printf("filenameOut: [%s]\n", filenameOut)
	}

	bytesInfo, err := os.ReadFile(filenameInfo)
	if err != nil {
		log.Fatalln("Failed to read file", filenameInfo, err)
	}

	bytesKeymap, err := os.ReadFile(filenameKeymap)
	if err != nil {
		log.Fatalln("Failed to read file", filenameKeymap, err)
	}

	qmkInfo := QMKInfo{}
	if err := json.Unmarshal(bytesInfo, &qmkInfo); err != nil {
		log.Fatalln("Failed to json unmarshal", filenameInfo, err)
	}

	qmkKeymap := QMKKeymap{}
	if err := json.Unmarshal(bytesKeymap, &qmkKeymap); err != nil {
		log.Fatalln("Failed to json unmarshal", filenameKeymap, err)
	}

	result := Draw(
		qmkInfo.Layouts,
		qmkKeymap,
	)
	if debug {
		log.Printf("Result:\n%s\n", result)
	}

	// Optional
	if filenameOut != "" {
		if err := os.WriteFile(filenameOut, []byte(result), 0o600); err != nil {
			log.Fatalln("Failed to write file", filenameOut, err)
		}
	}
}
