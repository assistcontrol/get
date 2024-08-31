package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

// Simple check for an HTTP-like protocol
var URLPattern = regexp.MustCompile(`^https?://`)

func fetch(url string) ([]byte, error) {
	// Try local file first
	contents, err := local(url)
	if err == nil {
		return contents, err
	}

	// Not a local file; fetch remotely
	contents, err = remote(url)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return contents, nil
}

func local(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func remote(url string) ([]byte, error) {
	if URLPattern.MatchString(url) {
		contents, err := getHTTP(url)
		if err != nil {
			return nil, err
		}
		return contents, nil
	}

	contents, err := getHTTP("https://" + url)
	if err != nil {
		contents, err = getHTTP("http://" + url)
	}
	if err != nil {
		return nil, err
	}

	return contents, nil
}
func getHTTP(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
