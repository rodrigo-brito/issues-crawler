package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rodrigo-brito/issues-crawler/crawler"
)

func main() {
	ctx := context.Background()
	fetcher := crawler.NewCrawler(ctx)
	resp, err := fetcher.Fetch("rodrigo-brito", "gocity")
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range resp {
		fmt.Printf("%#v", node)
	}
}
