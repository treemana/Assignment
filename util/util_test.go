package util

import (
	"net/url"
	"testing"
)

func TestGetPathNameFromURL(t *testing.T) {
	tests := []struct {
		name   string
		rawURL string
		want   string
		want1  string
	}{
		{name: "default", rawURL: "https://www.google.com", want: "", want1: "www.google.com.html"},
		{name: "no ext", rawURL: "https://www.google.com/a", want: "www.google.com", want1: "a.html"},
		{name: "full", rawURL: "https://www.google.com/b.c", want: "www.google.com", want1: "b.c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := url.Parse(tt.rawURL)
			if err != nil {
				t.Errorf("%s rawURL = %s is invalid", tt.name, tt.rawURL)
			}
			got, got1 := GetPathNameFromURL(u)
			if got != tt.want {
				t.Errorf("GetPathNameFromURL() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetPathNameFromURL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
