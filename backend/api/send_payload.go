package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func SendPayload(raw json.RawMessage, expire int, reads int) {
	payload := map[string]interface{}{
		"data":  raw,
		"ttl":   expire,
		"reads": reads,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal payload: %v\n", err)
		os.Exit(1)
	}

	resp, err := http.Post("http://localhost:8080/v1/payload", "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Unexpected status code: %d\n", resp.StatusCode)
		os.Exit(1)
	}

	fmt.Printf("Response status: %s\n", resp.Status)
	io.Copy(os.Stdout, resp.Body)
}