package util

import (
	"net/url"
	"path"
)

const (
	defaultSuffix = ".html"
)

// GetPathNameFromURL return the dir and filename located on local storage
// details can be found from test case
func GetPathNameFromURL(u *url.URL) (string, string) {

	if u == nil {
		return "", ""
	}

	if len(u.Path) == 0 {
		return "", u.Hostname() + defaultSuffix
	}

	dir, file := path.Split(u.Path)

	if len(path.Ext(file)) == 0 {
		file += defaultSuffix
	}

	return path.Join(u.Hostname(), dir), file
}
