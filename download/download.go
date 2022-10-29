package download

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/treemana/assignment/util"
)

// GetURLs download urls concurrently, one url one Goroutine
// if data file with same name already exist, the older one will be replaced by the newer one
func GetURLs(urls []*url.URL) {

	var resultChan = make(chan error, len(urls))

	for i := range urls {
		go func(index int) {
			if err := GetURL(urls[index]); err != nil {
				resultChan <- fmt.Errorf("download %s error, %s", urls[index].String(), err)
			}
			resultChan <- nil
		}(i)
	}

	for i := 0; i < len(urls); i++ {
		if err := <-resultChan; err != nil {
			fmt.Println(err)
		}
	}
}

// GetURL download url data to local disk
// if data file with same name already exist, the older one will be replaced by the newer one
func GetURL(u *url.URL) error {
	dir, name := util.GetPathNameFromURL(u)
	if len(dir) != 0 {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("[%s]", dir)
			return err
		}
	}

	pathname := filepath.Join(dir, name)
	old, err := os.Stat(pathname)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if old != nil && old.IsDir() {
		return errors.New("storage path already exist as directory")
	}

	// prepare data source
	var resp *http.Response
	if resp, err = http.Get(u.String()); err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	// prepare data target
	var f *os.File
	if f, err = os.Create(pathname); err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	// copy data from source to target
	if _, err = io.Copy(f, resp.Body); err != nil {
		return err
	}

	return nil
}
