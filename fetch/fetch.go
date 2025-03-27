// Package fetch provides a functions to fetch from a local path or a remote URL.
package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/assistcontrol/get/context"
)

// Simple check for an HTTP-like protocol
var urlPattern = regexp.MustCompile(`^https?://`)

// Fetch retrieves the contents of a file from a local path or a remote URL.
// It first tries to read the file from the local filesystem. If that fails,
// it tries to fetch the file from the remote URL. If that fails, it returns
// an error.
func Fetch(c *context.Ctx) (err error) {
	// Try local file first
	if err = local(c); err == nil {
		setLocalFilename(c)
		return
	}

	// Non-nil err means local file not found, so try fetching
	// it remotely
	if err = remote(c); err == nil {
		setRemoteFilename(c)
	}

	return
}

// local reads a file from the local filesystem.
// It returns an error if the file isn't readable or doesn't exist.
func local(c *context.Ctx) error {
	b, err := os.ReadFile(c.Path)
	c.Body = b

	return err
}

// remote oversees fetching a file from a remote URL.
func remote(c *context.Ctx) error {
	// If it's got a protocol already, fetch it directly
	if urlPattern.MatchString(c.Path) {
		c.URL = c.Path
		return getHTTP(c)
	}

	// Try HTTPS, then HTTP, then just give up
	c.URL = "https://" + c.Path
	err := getHTTP(c)
	if err != nil {
		c.URL = "http://" + c.Path
		err = getHTTP(c)
	}

	return err
}

// getHTTP fetches a file from a remote URL using the HTTP protocol.
// If successful, it saves the http.Response object and extracts the
// body of the response. It returns an error if the file cannot be
// fetched or something went wacky.
func getHTTP(c *context.Ctx) error {
	resp, err := http.Get(c.URL)
	switch {
	case err != nil:
		return err
	case resp.StatusCode != http.StatusOK:
		return fmt.Errorf("HTTP error: %s", resp.Status)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic("Response body close error: " + err.Error())
		}
	}()

	if c.Body, err = io.ReadAll(resp.Body); err != nil {
		return err
	}

	c.Response = resp

	return err
}
