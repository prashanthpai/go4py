package main

import (
	"fmt"
	"runtime"
)

func main() {
	// START OMIT
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Go is running on Linux.")
		// fallthrough
	case "darwin":
		fmt.Println("Go is running on OS X.")
	default:
		fmt.Printf("Go is running on %s\n", os)
	}
	// END OMIT
}
