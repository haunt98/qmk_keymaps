package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	filenameInfo := os.Getenv("QMK_INFO")
	if filenameInfo == "" {
		log.Fatalln("ENV QMK_INFO empty")
	}

	filenameKeymap := os.Getenv("QMK_KEYMAP")
	if filenameKeymap == "" {
		log.Fatalln("ENV QMK_KEYMAP empty")
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
	fmt.Println(result)
}
