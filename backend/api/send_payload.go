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
	if expire <= 0 {
		expire = 10
	}

	payload := struct {
		Data           json.RawMessage `json:"data"`
		TTL            int            `json:"ttl"`
		RemainingReads int            `json:"remaining_reads"`
	}{
		Data:           raw,
		TTL:            expire,
		RemainingReads: reads,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal payload: %v\n", err)
		os.Exit(1)
	}

	req, err := http.NewRequest(http.MethodPost, "https://restapi.egeuysal.com/v1/payload", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create request: %v\n", err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Origin", "https://restapi.egeuysal.com")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read response: %v\n", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Error: %s\nResponse: %s\n", resp.Status, string(body))
		os.Exit(1)
	}

	fmt.Printf("Response status: %s\n", resp.Status)
	fmt.Printf("Response body: %s\n", string(body))
}