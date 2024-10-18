package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %s\n", os.Args[1])
	html, err := getHTML(os.Args[1])
	if err != nil {
		fmt.Printf("error fetching html: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(html)
}
