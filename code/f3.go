package main

import (
    "fmt"
    "time"
)

func main() {
// START OMIT
    i := 0
	for {
        i++
        fmt.Printf("%d\n", i)
        time.Sleep(100 * time.Millisecond)
	}
// END OMIT
}
