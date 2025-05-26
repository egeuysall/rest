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

const version = "1.0.0"

func GetUsedFlags() map[string]bool {
	used := map[string]bool{
		"data": false, "d": false,
		"expire": false, "e": false,
		"once": false, "o": false,
		"help": false, "h": false,
		"version": false, "v": false,
	}

	for _, arg := range os.Args[1:] {
		if len(arg) > 1 && arg[0] == '-' {
			argName := arg
			if arg[1] == '-' && len(arg) > 2 {
				argName = arg[2:]
			} else {
				argName = arg[1:]
			}
			if idx := indexOf(argName, '='); idx != -1 {
				argName = argName[:idx]
			}
			if _, ok := used[argName]; ok {
				used[argName] = true
			}
		}
	}

	return used
}

func indexOf(s string, sep byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == sep {
			return i
		}
	}
	return -1
}

func RandomSlug(n int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, n)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result[i] = chars[num.Int64()]
	}
	return string(result)
}

var helpMessage = `
Usage:
  rest [options]

Options:
  -data, --d           JSON string to upload or @file.json to load
  -expire, --e         Minutes until the endpoint expires (default: 10)
  -once, --o           Delete endpoint after first access
  -version, --v        Show version information
  -help, --h           Show this help message

Examples:
  rest --d payload.json --o
`

type Config struct {
	Data    string
	Expire  int
	Once    bool
	Help    bool
	Version bool
}

func bindFlags(config *Config) {
	flag.StringVar(&config.Data, "data", config.Data, "Data string to process")
	flag.StringVar(&config.Data, "d", config.Data, "Data string to process (shorthand)")

	flag.IntVar(&config.Expire, "expire", config.Expire, "Expiration time in minutes (default: 10)")
	flag.IntVar(&config.Expire, "e", config.Expire, "Expiration time in minutes (shorthand)")

	flag.BoolVar(&config.Once, "once", config.Once, "Process only once (auto-delete after first use)")
	flag.BoolVar(&config.Once, "o", config.Once, "Process only once (shorthand)")

	flag.BoolVar(&config.Help, "help", config.Help, "Show help message")
	flag.BoolVar(&config.Help, "h", config.Help, "Show help message (shorthand)")

	flag.BoolVar(&config.Version, "version", config.Version, "Show version information")
	flag.BoolVar(&config.Version, "v", config.Version, "Show version information (shorthand)")

	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), helpMessage)
	}
}

func handleSpecialFlags(config *Config, usedFlags map[string]bool) {
	if usedFlags["once"] || usedFlags["o"] {
		fmt.Println("Custom help logic triggered!")
		os.Exit(0)
	}

	if config.Help {
		flag.Usage()
		os.Exit(0)
	}

	if config.Version {
		fmt.Println("Version", version)
		os.Exit(0)
	}
}

func ParseFlags() *Config {
	config := &Config{
		Data:    "",
		Expire:  10,
		Once:    false,
		Help:    false,
		Version: false,
	}

	bindFlags(config)
	flag.Parse()

	usedFlags := GetUsedFlags()
	for k, v := range usedFlags {
		if v {
			fmt.Println("Used flag:", k)
		}
	}

	handleSpecialFlags(config, usedFlags)

	return config
}

func validateJSON(data []byte) error {
	if !json.Valid(data) {
		return fmt.Errorf("invalid JSON format")
	}
	return nil
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

	if err := validateJSON(data); err != nil {
		return nil, err
	}

	return json.RawMessage(data), nil
}