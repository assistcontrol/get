package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

// Simple check for an HTTP-like protocol
var URLPattern = regexp.MustCompile(`^https?://`)

func main() {
	args := os.Args
	if len(args) < 2 {
		help()
		return
	}
	url := args[1]

	contents, err := os.ReadFile(url)
	if err != nil {
		contents, err = get(url)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	show(contents)
}

func help() {
	fmt.Println("Usage: get <url>")
}

func get(url string) ([]byte, error) {
	if URLPattern.MatchString(url) {
		contents, err := fetch(url)
		if err != nil {
			return nil, err
		}
		return contents, nil
	}

	contents, err := fetch("https://" + url)
	if err != nil {
		contents, err = fetch("http://" + url)
	}
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
