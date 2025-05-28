package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  rest server    - Start the server")
		fmt.Println("  rest client    - Run as client (use with -d, -t, -e flags)")
		fmt.Println("\nClient flags:")
		fmt.Println("  -d, -data      JSON string or file path (required)")
		fmt.Println("  -t, -times     Delete after N accesses (0 = infinite, default: 1)")
		fmt.Println("  -e, -expire    Minutes until expiration (default: 10)")
		fmt.Println("  -h, -help      Show help message")
		fmt.Println("  -v, -version   Show version")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "server":
		os.Args = append([]string{os.Args[0]}, os.Args[2:]...)
		main()
	case "client":
		os.Args = append([]string{os.Args[0]}, os.Args[2:]...)
		main()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}