package main

import (
	"flag"
	"fmt"
)

func main() {
	path := flag.String("p", "", "Path to a root dir")
	flag.Parse()
	extensions = flag.Args()

	p := getPath(*path)
	fmt.Println("\nRoot dir: ", p)
	readDirectory(p)
	fmt.Println("Total:", numberOfLines, "lines")
}
