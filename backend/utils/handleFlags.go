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
	
	fileData, err := ReadJSONFile(data)
	if err != nil {
		return nil, err
	}
	
	if !json.Valid(fileData) {
		return nil, fmt.Errorf("invalid JSON format")
	}
	
	return fileData, nil
}

func handleSpecialFlags(config *Config) {
	if config.Data != "" {
		data, err := handleDataFlag(config.Data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		config.RawData = data
	}

	switch {
	case config.Times == 0:
		fmt.Println("Debug: Times set to infinite")
	case config.Help:
		flag.Usage()
		os.Exit(0)
	case config.Version:
		fmt.Printf("Version %s\n", version)
		os.Exit(0)
	}
}