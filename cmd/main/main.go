package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ekeydar/link"
)

func main() {
	filename := flag.String("file", "ex1.html", "the HTML file to parse")
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Failed to open file %s: %s", *filename, err)
	}
	links, parseErr := link.ParseFile(f)
	if parseErr != nil {
		log.Fatalf("Failed to parse %s: %s", *filename, err)
	}
	for _, link := range links {
		fmt.Printf("href=%s\n%s\n-------------\n", link.Href, link.Text)
	}
}
