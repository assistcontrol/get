package filename

import (
	"mime"
	"net/http"
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

func FromLocalPath(path string) string {
	return filepath.Base(path)
}

func FromHTTPResponse(resp *http.Response) string {
	basename := filepath.Base(resp.Request.URL.RequestURI())

	// If we have a complete filename, return it
	if filenameWithExtension.MatchString(basename) {
		return basename
	}

	// Determine an appropriate extension. If we can't, return the default
	mimetype := resp.Header.Get("Content-Type")
	if mimetype == "" {
		return defaultFilename
	}
	extensions, err := mime.ExtensionsByType(mimetype)
	if err != nil || len(extensions) == 0 {
		return defaultFilename
	}
	extension := extensions[len(extensions)-1]

	// See if we can deduce a name, otherwise make up a filename
	if filenameWithoutExtension.MatchString(basename) {
		return basename + extension
	}
	return defaultFilename + extension
}
