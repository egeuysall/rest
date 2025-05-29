package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	filePath := flag.String("d", "", "Path to JSON file")
	times := flag.Int("t", 1, "Number of times the payload can be accessed")
	expiresIn := flag.Int("e", 0, "Expires in minutes")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide a JSON file path using -d flag")
		os.Exit(1)
	}

	jsonData, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var jsonObj interface{}
	if err := json.Unmarshal(jsonData, &jsonObj); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	payload := map[string]interface{}{
		"data": jsonObj,
	}

	if *times > 0 {
		payload["remaining_reads"] = *times
	}

	if *expiresIn > 0 {
		payload["expires_in"] = *expiresIn
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", "https://restapi.egeuysal.com/v1/payload", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(string(body))
		os.Exit(1)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Payload stored successfully!\n")
	fmt.Printf("Access URL: %s\n", result["url"])
	fmt.Printf("ID: %s\n", result["id"])
} 