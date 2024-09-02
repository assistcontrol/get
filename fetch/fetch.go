package fetch

import (
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/assistcontrol/get/context"
)

// Simple check for an HTTP-like protocol
var urlPattern = regexp.MustCompile(`^https?://`)

func Fetch(c *context.Ctx) error {
	// Try local file first
	if err := local(c); err == nil {
		c.SetLocalFilename()
		return nil
	}

	if err := remote(c); err != nil {
		return err
	}

	c.SetRemoteFilename()
	return nil
}

func local(c *context.Ctx) error {
	b, err := os.ReadFile(c.Path)
	c.Body = b

	return err
}

func remote(c *context.Ctx) error {
	if urlPattern.MatchString(c.Path) {
		// It's a full URL, so fetch it directly
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

func getHTTP(c *context.Ctx) error {
	resp, err := http.Get(c.URL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	c.Body, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	c.Response = resp

	return err
}
