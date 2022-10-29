package metadata

import (
	"fmt"
	"time"
)

const (
	timeLayout = "Mon Jan 02 2006 15:04 MST"
)

type Meta struct {
	Error     error
	Site      string
	NumLinks  uint
	Images    uint
	LastFetch *time.Time
}

func (m *Meta) Print() {
	fmt.Println()
	fmt.Println("site:", m.Site)
	fmt.Println("num_links:", m.NumLinks)
	fmt.Println("images:", m.Images)
	if m.LastFetch == nil {
		fmt.Println("last_fetch: not downloaded yet")
	} else {
		fmt.Println("last_fetch:", m.LastFetch.Format(timeLayout))
	}

	fmt.Println()
}
