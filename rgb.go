package main

import (
	"fmt"
)

func main() {

	fmt.Print("Type an RGB: ")
	var r, g, b int

	fmt.Scan(&r, &g, &b)

	fmt.Printf("#%02x%02x%02x", r,g,b)
}