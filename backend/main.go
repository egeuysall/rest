package main

import (
	"fmt"

	"github.com/egeuysall/rest/utils"
)

func main() {
    config := utils.ParseFlags()
    used := utils.GetUsedFlags()

    if used["d"] || used["data"] {
        fmt.Println("Data flag used")
    }

    if used["o"] || used["once"] {
        fmt.Println("Once flag used")
    }

    fmt.Printf("Data: %s\n", config.Data)
}