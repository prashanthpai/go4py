package main

import "fmt"

// START OMIT
func generateEven(count int, channel chan int) {
	for i := 0; count > 0; i++ {
		if i%2 == 0 {
			channel <- i    // yield
			count--
            // fmt.Printf("generateEven() sent: %d\n", i)
		}
	}
	close(channel)          // StopIteration
}

func main() {
	channel := make(chan int)
	go generateEven(3, channel)

	for number := range channel {
        fmt.Printf("main() received: %d\n", number)
	}
}
// END OMIT
