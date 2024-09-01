package fetch

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/assistcontrol/get/body"
	"github.com/assistcontrol/get/config"
	"github.com/assistcontrol/get/fetch/filename"
)

const defaultFilename = "get.output"

var (

	// Simple check for an HTTP-like protocol
	urlPattern = regexp.MustCompile(`^https?://`)
)

func Fetch(c *config.Config) (*body.Body, error) {
	// Try local file first
	b, err := local(c.URL)
	if err == nil {
		return b, err
	}

	// Not a local file; fetch remotely
	b, err = remote(c.URL)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return b, nil
}

func local(path string) (*body.Body, error) {
	contents, err := os.ReadFile(path)

	b := &body.Body{
		Body:     contents,
		Filename: filename.FromLocalPath(path),
	}

	return b, err
}

func remote(url string) (b *body.Body, err error) {
	switch urlPattern.MatchString(url) {
	case true:
		// It's a full URL, so fetch it directly
		b, err = getHTTP(url)

	case false:
		// Try HTTPS, then HTTP, then just give up
		b, err = getHTTP("https://" + url)
		if err != nil {
			b, err = getHTTP("http://" + url)
		}
	}

	return
}

func getHTTP(url string) (*body.Body, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	b := &body.Body{
		Body:     contents,
		Filename: filename.FromHTTPResponse(resp),
	}

	return b, err
}
