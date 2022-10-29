package metadata

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/net/html"

	"github.com/treemana/assignment/util"
)

func PrintFromURLs(urls []*url.URL) {
	var resultChan = make(chan *Meta, len(urls))

	for i := range urls {
		go func(index int) {
			meta, err := GetFromURL(urls[index])
			if err != nil {
				fmt.Printf("get %s metadata error, %s", urls[index].String(), err)
			}
			resultChan <- meta
		}(i)
	}

	var metadataList = make([]*Meta, 0, len(urls))
	for i := 0; i < len(urls); i++ {
		if meta := <-resultChan; meta != nil {
			metadataList = append(metadataList, meta)
		}
	}

	if len(metadataList) != len(urls) {
		return
	}

	for _, meta := range metadataList {
		meta.Print()
	}
}

func GetFromURL(u *url.URL) (*Meta, error) {
	dir, name := util.GetPathNameFromURL(u)
	pathname := filepath.Join(dir, name)
	fi, err := os.Stat(pathname)
	if err != nil {
		if os.IsNotExist(err) {
			return &Meta{Site: u.String(), LastFetch: nil}, nil
		}
		return nil, err
	}

	lastFetch := fi.ModTime()
	var meta = &Meta{
		Site:      u.String(),
		LastFetch: &lastFetch,
	}

	var f *os.File
	if f, err = os.Open(pathname); err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	var root *html.Node
	if root, err = html.Parse(f); err != nil {
		return nil, err
	}

	// level order traversal of html.Node tree
	var nodes = []*html.Node{root, nil}
	var newLevel bool
	for i := 0; i < len(nodes); i++ {
		// end of one level
		if nodes[i] == nil {
			// next level exist
			if newLevel {
				newLevel = false
				nodes = append(nodes, nil)
				continue
			}
			// next level not exist
			break
		}
		node := nodes[i]
		if node.Type == html.ElementNode {
			switch node.Data {
			case "a", "area", "link":
				for _, element := range node.Attr {
					if element.Key == "href" {
						meta.NumLinks++
					}
				}
			case "audio", "iframe":
				for _, element := range node.Attr {
					if element.Key == "src" {
						meta.NumLinks++
					}
				}
			case "command":
				for _, element := range node.Attr {
					if element.Key == "icon" {
						meta.Images++
					}
				}
			case "img", "picture":
				for _, element := range node.Attr {
					if element.Key == "src" {
						meta.Images++
					}
				}
			}
		}

		// push all children into queue only if exist
		child := node.FirstChild
		if child != nil {
			newLevel = true
			for child != nil {
				nodes = append(nodes, child)
				child = child.NextSibling
			}
		}
	}

	return meta, nil
}
