package main

import (
	"fmt"

	"github.com/egeuysall/rest/utils"
)

func main() {
	config := utils.ParseFlags()
	fmt.Printf("JSON Data:\n%s\n", config.RawData)
}