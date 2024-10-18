package main

import (
	"fmt"
	"sort"
)

type Page struct {
	URL   string
	Count int
}

func printReport(pages map[string]int, baseURL string) {
	printHeader(baseURL)

	sorted := sortPage(pages)

	for _, entry := range sorted {
		fmt.Printf("Found %d internal links to %s\n", entry.Count, entry.URL)
	}
}

func printHeader(baseURL string) {
	fmt.Printf("\n")
	for i := 0; i < len(baseURL)+11; i++ {
		fmt.Printf("=")
	}
	fmt.Printf("\n")
	fmt.Println("REPORT for", baseURL)
	for i := 0; i < len(baseURL)+11; i++ {
		fmt.Printf("=")
	}
	fmt.Printf("\n")
}

func sortPage(pages map[string]int) []Page {
	sorted := []Page{}

	for normalizedURL, count := range pages {
		sorted = append(sorted, Page{
			URL:   normalizedURL,
			Count: count,
		})
	}

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Count == sorted[j].Count {
			return sorted[i].URL < sorted[j].URL
		}
		return sorted[i].Count > sorted[j].Count
	})
	return sorted
}
