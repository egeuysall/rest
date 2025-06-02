package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func handleDataFlag(data string) (json.RawMessage, error) {
	if data == "" {
		return nil, fmt.Errorf("no data provided")
	}
	
	fileData, err := ReadJsonFile(data)
	if err != nil {
		return nil, err
	}
	
	if !json.Valid(fileData) {
		return nil, fmt.Errorf("invalid JSON format")
	}
	
	return fileData, nil
}

func HandleSpecialFlags(config *Config) {
	if config.Data != "" {
		data, err := handleDataFlag(config.Data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		config.RawData = data
	}

	if config.Help {
		flag.Usage()
		os.Exit(0)
	}

	if config.Times >= -1 {
		if config.Times == -1 {
			fmt.Println("Times set to infinite")
		} else {
			fmt.Printf("Times set to %d\n", config.Times)
		}
	}

	if config.Expire >= -1 {
		if config.Expire == -1 {
			fmt.Println("Expire set to infinite")
		} else {
			fmt.Printf("Expire set to %d\n", config.Expire)
		}
	}

	// If RawData is set, send the payload
	if config.RawData != nil {
		SendPayload(config.RawData, config.Expire, config.Times)
	}
}