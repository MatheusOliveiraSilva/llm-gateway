// Package main is a simple hello world application that greets users by name.
package main

import (
	"flag"
	"fmt"
)

func main() {

	name := flag.String("name", "Assistant", "The name to greet")
	flag.Parse()

	fmt.Printf("Oi, %s! 🚀", *name)
}
