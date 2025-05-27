package utils

import (
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"
)

var version = "1.0.0"

type Config struct {
	Data    string
	RawData json.RawMessage
	Expire  int
	Once    bool
	Help    bool
	Version bool
}

func ParseFlags() *Config {
	config := &Config{
		Expire: 10,
	}

	// Define flags
	flag.StringVar(&config.Data, "data", "", "JSON string or file path")
	flag.StringVar(&config.Data, "d", "", "JSON string or file path (shorthand)")
	flag.IntVar(&config.Expire, "expire", 10, "Expiration time in minutes")
	flag.IntVar(&config.Expire, "e", 10, "Expiration time in minutes (shorthand)")
	flag.BoolVar(&config.Once, "once", false, "Delete after first access")
	flag.BoolVar(&config.Once, "o", false, "Delete after first access (shorthand)")
	flag.BoolVar(&config.Help, "help", false, "Show help message")
	flag.BoolVar(&config.Help, "h", false, "Show help message (shorthand)")
	flag.BoolVar(&config.Version, "version", false, "Show version")
	flag.BoolVar(&config.Version, "v", false, "Show version (shorthand)")

	// Custom help message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), `
Usage: rest [options]

Options:
  -data, -d     JSON string or file path (required)
  -expire, -e   Minutes until expiration (default: 10)
  -once, -o     Delete after first access
  -version, -v  Show version
  -help, -h     Show this help message

Example: rest -d payload.json -o
`)
	}

	flag.Parse()
	
	if config.Data == "" {
		fmt.Fprintf(os.Stderr, "Error: -data or -d flag is required\n")
		flag.Usage()
		os.Exit(1)
	}

	handleSpecialFlags(config)
	return config
}

func ReadJSONFile(filePath string) (json.RawMessage, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	
	return data, nil
}

func randomSlug(n int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, n)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result[i] = chars[num.Int64()]
	}
	return string(result)
}