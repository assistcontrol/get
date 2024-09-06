package fetch

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/assistcontrol/get/context"
)

var localScenarios = []struct {
	expected string
	ctx      *context.Ctx
}{
	{
		expected: "output.txt",
		ctx: &context.Ctx{
			Path: "output.txt",
		},
	},
	{
		expected: "output.txt",
		ctx: &context.Ctx{
			Path: "/foo/bar/output.txt",
		},
	},
	{
		expected: "output.txt",
		ctx: &context.Ctx{
			Path: "./foo/output.txt",
		},
	},
	{
		expected: "foo",
		ctx: &context.Ctx{
			Path:     "output.txt",
			Filename: "foo",
		},
	},
}

func TestSetLocalFilename(t *testing.T) {
	for _, s := range localScenarios {
		setLocalFilename(s.ctx)
		if s.ctx.Destination != s.expected {
			t.Errorf("expected %q, got %q", s.expected, s.ctx.Destination)
		}
	}
}

var remoteScenarios = []struct {
	url      string
	expected string
	cType    string
}{
	{
		url:      "https://example.com",
		expected: "get.output",
		cType:    "",
	},
	{
		url:      "https://example.com",
		expected: "get.output.html",
		cType:    "text/html",
	},
	{
		url:      "https://example.com/foo",
		expected: "foo.html",
		cType:    "text/html",
	},
	{
		url:      "https://example.com/foo",
		expected: "get.output",
		cType:    "",
	},
	{
		url:      "https://example.com/foobar.txt",
		expected: "foobar.txt",
		cType:    "text/html",
	},
	{
		url:      "https://example.com",
		expected: "get.output.gif",
		cType:    "image/gif",
	},
}

func TestSetRemoteFilename(t *testing.T) {
	for _, s := range remoteScenarios {
		u, _ := url.Parse(s.url)

		ctx := &context.Ctx{
			Response: &http.Response{
				Header: http.Header{},
				Request: &http.Request{
					URL: u,
				},
			},
		}
		ctx.Response.Header.Set("Content-Type", s.cType)

		setRemoteFilename(ctx)
		if ctx.Destination != s.expected {
			t.Errorf("expected %q, got %q", s.expected, ctx.Destination)
		}
	}
}
