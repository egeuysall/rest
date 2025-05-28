package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var version string = "1.0.0"

type Config struct {
	Data    string
	RawData json.RawMessage
	Expire  int
	Times   int
	Help    bool
	Version bool
}

func ParseFlags() *Config {
	config := &Config{
		Expire: 10,
		Times:  1,
	}

	flag.StringVar(&config.Data, "data", "", "JSON string or file path")
	flag.StringVar(&config.Data, "d", "", "JSON string or file path (shorthand)")
	flag.IntVar(&config.Expire, "expire", 10, "Expiration time in minutes (default: 10)")
	flag.IntVar(&config.Expire, "e", 10, "Expiration time in minutes (default: 10)")
	flag.IntVar(&config.Times, "times", 1, "Number of times the data can be accessed before deletion")
	flag.IntVar(&config.Times, "t", 1, "Number of times the data can be accessed before deletion (shorthand)")
	flag.BoolVar(&config.Help, "help", false, "Show help message")
	flag.BoolVar(&config.Help, "h", false, "Show help message (shorthand)")
	flag.BoolVar(&config.Version, "version", false, "Show version")
	flag.BoolVar(&config.Version, "v", false, "Show version (shorthand)")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), `
Usage: rest [options]

Options:
  -data, -d     JSON string or file path (required)
  -expire, -e   Minutes until expiration (default: 10)
  -times, -t    Delete after N accesses (0 = infinite, default: 1)
  -version, -v  Show version
  -help, -h     Show this help message

Example: rest -d payload.json -t 5 -e 30
`)
	}

	flag.Parse()

	if config.Help {
		flag.Usage()
		os.Exit(0)
	}

	if config.Version {
		fmt.Printf("Version %s\n", version)
		os.Exit(0)
	}

	if config.Data == "" {
		fmt.Fprintln(os.Stderr, "Error: -data or -d flag is required")
		flag.Usage()
		os.Exit(1)
	}

	if config.Expire < 1 {
		config.Expire = 10
	}

	var raw json.RawMessage
	if strings.HasSuffix(config.Data, ".json") || strings.HasSuffix(config.Data, ".txt") {
		r, err := ReadJSONFile(config.Data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
			os.Exit(1)
		}
		raw = r
	} else {
		raw = json.RawMessage(config.Data)
	}

	config.RawData = raw

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

	if !json.Valid(data) {
		return nil, fmt.Errorf("invalid JSON in file: %s", string(data))
	}

	return data, nil
}