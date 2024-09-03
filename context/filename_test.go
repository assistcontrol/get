package context

import "testing"

var localScenarios = []struct {
	expected string
	ctx      *Ctx
}{
	{
		expected: "output.txt",
		ctx: &Ctx{
			Path: "output.txt",
		},
	},
	{
		expected: "output.txt",
		ctx: &Ctx{
			Path: "/foo/bar/output.txt",
		},
	},
	{
		expected: "output.txt",
		ctx: &Ctx{
			Path: "./foo/output.txt",
		},
	},
	{
		expected: "foo",
		ctx: &Ctx{
			Path:     "output.txt",
			Filename: "foo",
		},
	},
}

func TestSetLocalFilename(t *testing.T) {
	for _, s := range localScenarios {
		s.ctx.SetLocalFilename()
		if s.ctx.Destination != s.expected {
			t.Errorf("expected %q, got %q", s.expected, s.ctx.Destination)
		}
	}
}
