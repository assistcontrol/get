package context

import (
	"mime"
	"path/filepath"
	"regexp"
)

const defaultFilename = "get.output"

var (
	// Simple check for a filename with an extension
	filenameWithExtension = regexp.MustCompile(`\.\w+$`)
	// Simple check for a filename without an extension
	filenameWithoutExtension = regexp.MustCompile(`^\w+$`)
)

func (c *Ctx) SetLocalFilename() {
	if c.Filename != "" {
		c.Destination = c.Filename
		return
	}

	c.Destination = filepath.Base(c.Path)
}

func (c *Ctx) SetRemoteFilename() {
	if c.Filename != "" {
		c.Destination = c.Filename
		return
	}

	if c.Response == nil {
		c.Destination = defaultFilename
		return
	}

	basename := filepath.Base(c.Response.Request.URL.RequestURI())

	// If we have a complete filename, use it
	if filenameWithExtension.MatchString(basename) {
		c.Destination = basename
		return
	}

	// Determine an appropriate extension. If we can't, return the default
	mimetype := c.Response.Header.Get("Content-Type")
	if mimetype == "" {
		c.Destination = defaultFilename
		return
	}
	extensions, err := mime.ExtensionsByType(mimetype)
	if err != nil || len(extensions) == 0 {
		c.Destination = defaultFilename
		return
	}
	extension := extensions[len(extensions)-1]

	// See if we can deduce a name, otherwise make up a filename
	if filenameWithoutExtension.MatchString(basename) {
		c.Destination = basename + extension
		return
	}

	c.Destination = defaultFilename + extension
}
