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
