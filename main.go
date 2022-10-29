package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/treemana/assignment/download"
	"github.com/treemana/assignment/metadata"
)

const (
	metadataKey = "--metadata"
)

var (
	metadataMap map[string]struct{}
)

func init() {
	metadataMap = map[string]struct{}{metadataKey: {}}
}

func main() {
	fmt.Println("fetch starting ...")

	var urlMap = make(map[string]*url.URL, len(os.Args)-1)
	var metadataEnable bool

	// gather command line parameters
	for _, arg := range os.Args[1:] {
		if _, ok := metadataMap[arg]; ok {
			metadataEnable = true
			continue
		}

		u, err := url.ParseRequestURI(arg)
		if err != nil {
			fmt.Printf("[%s] is invalid as a web site link\n", arg)
			return
		}

		urlMap[u.String()] = u
	}

	if len(urlMap) == 0 {
		fmt.Println("please input at least one link")
		return
	}

	var urls = make([]*url.URL, 0, len(urlMap))
	for _, u := range urlMap {
		urls = append(urls, u)
	}

	// print metadata
	if metadataEnable {
		metadata.PrintFromURLs(urls)
		return
	}

	// download from origin
	download.GetURLs(urls)

	fmt.Println("fetch stopped")

}
