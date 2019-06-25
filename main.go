package main

import (
	"fmt"
)

// HelloWorld hello world
func HelloWorld(s string) string {
	return "Hello World, " + s
}

func main() {
	fmt.Println("main")
	HelloWorld("n0ts")
}
