package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadJsonFile(filePath string) (json.RawMessage, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if !json.Valid(data) {
		return nil, fmt.Errorf("invalid JSON in file: %s", string(data))
	}

	return data, nil
}