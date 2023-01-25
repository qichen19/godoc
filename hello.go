package main

import (
	"example/hello/dog"
	"fmt"
)

// main function
func main() {
	y := 3
	dy := dog.Years(y)
	fmt.Printf("%v human year is equal to %v dog year", y, dy)
}
