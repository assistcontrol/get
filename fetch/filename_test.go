package fetch

import (
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
