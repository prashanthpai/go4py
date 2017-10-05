package main

import "fmt"

// START OMIT
func stringInSlice(query string, list []string) bool {
	for _, s := range list {
		if s == query {
			return true
		}
	}
	return false
}

func main() {
    list := []string{"gluster", "go", "python"}
    fmt.Printf("%t\n", stringInSlice("go", list))
    fmt.Printf("%t\n", stringInSlice("java", list))
}
// END OMIT
