package main

import (
	"fmt"
	"time"

	"github.com/egeuysall/rest/api"
	"github.com/egeuysall/rest/db"
)

func main() {
	db.Connect()
	defer db.Close()
	
	go func() {
		ticker := time.NewTicker(3 * time.Minute)
		defer ticker.Stop()
		for {
			<-ticker.C
			if err := db.DeleteExpiredPayloads(); err != nil {
				fmt.Printf("Error cleaning up expired payloads: %v\n", err)
			} else {
				fmt.Println("Expired payloads cleaned up.")
			}
		}
	}()

	api.StartServer()

	fmt.Println("Server running on https://restapi.egeuysal.com/v1/payload")
} 