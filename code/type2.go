package main

import "fmt"

// START OMIT
func add(a, b int) int {
    return a + b
}

func multiply(a, b int) int {
    return a * b
}

func main() {
    fmt.Println(add(2, 3))
    fmt.Println(multiply(2, 3))
}
// END OMIT
