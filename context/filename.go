package context

import (
	"mime"
	"path/filepath"
	"regexp"
)

// defaultFilename is the filename used when the user hasn't specified one,
// or if we can't determine one. Rather than being useful, it proritizes
// making it clear where the file came from.
const defaultFilename = "get.output"

var (
	baseWithExtension    = regexp.MustCompile(`\.\w+$`) // A filename that includes an extension
	baseWithoutExtension = regexp.MustCompile(`^\w+$`)  // A filename with no extension
)

// SetLocalFilename sets the destination filename for a local file.
// If the user has specified a filename, it uses that.
func (c *Ctx) SetLocalFilename() {
	if c.Filename != "" {
		c.Destination = c.Filename
		return
	}

	c.Destination = filepath.Base(c.Path)
}

// SetRemoteFilename sets the destination filename for a remote file.
// If the user has specified a filename, it uses that.
// It tries to determine the filename from the URL and Content-Type.
// If it can't, it uses a default filename. This function should be
// called after the response has been received, or it will fall back
// to the default filename.
func (c *Ctx) SetRemoteFilename() {
	// If the user has specified a filename, use it.
	if c.Filename != "" {
		c.Destination = c.Filename
		return
	}

	// Handle a nil http.Response if that somehow happens.
	if c.Response == nil {
		c.Destination = defaultFilename
		return
	}

	// Get the basename of the final URL, i.e. post-redirs, etc.
	basename := filepath.Base(c.Response.Request.URL.RequestURI())

	// If we have a complete filename, use it
	if baseWithExtension.MatchString(basename) {
		c.Destination = basename
		return
	}

	// If we have a filename with no extension, we're looking at
	// something like example.com/ or example.com/foo.

	// Extract the mimetype from response headers.
	mimetype := c.Response.Header.Get("Content-Type")
	if mimetype == "" {
		c.Destination = defaultFilename
		return
	}

	// Go only includes extentions for a couple mimetypes, but it
	// should be enough for the situation we're in. It's pretty
	// unlikely that a .../foo or .../foo/ endpoint will be
	// anything other than text/html.
	extensions, err := mime.ExtensionsByType(mimetype)
	if err != nil || len(extensions) == 0 {
		c.Destination = defaultFilename
		return
	}

	// The mime package looks at system mimetype lists. Those lists
	// are almost always constructed with the most common extension
	// first, which mime puts last because reasons.
	extension := extensions[len(extensions)-1]

	// In the ../foo case, use "foo" for the filename.
	if baseWithoutExtension.MatchString(basename) {
		c.Destination = basename + extension
		return
	}

	// Without an obvious filename (.../foo/ or example.com), just
	// fall back to the default.
	c.Destination = defaultFilename + extension
}
