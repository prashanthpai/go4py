package main

import "fmt"

func main() {
// START OMIT
    var m map[string]int
    m = make(map[string]int)

    m["age"] = 33
    m["weight"] = 77

    for k, v := range m {
        fmt.Printf("%s %d\n", k, v)
    }

    delete(m, "age")
// END OMIT
}

