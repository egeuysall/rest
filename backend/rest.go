package main

import (
	"fmt"

	"github.com/egeuysall/rest/api"
	"github.com/egeuysall/rest/utils"
	"github.com/egeuysall/rest/db"
)

func main() {
	db.Connect()
	defer db.Close()

	config := utils.ParseFlags()
	fmt.Printf("JSON Data:\n%s\n", config.RawData)

	api.StartServer()
}