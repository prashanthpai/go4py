package main

import "fmt"

func main() {
// START OMIT
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
// END OMIT
}
