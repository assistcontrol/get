package context

import "testing"

var scenarios = []struct {
	expected string
	ctx      *Ctx
}{
	{
		expected: "output.txt",
		ctx: &Ctx{
			Filename: "output.txt",
		},
	},
	{
		expected: "output.txt",
		ctx: &Ctx{
			Filename: "/foo/bar/output.txt",
		},
	},
	{
		expected: "output.txt",
		ctx: &Ctx{
			Filename: "./foo/output.txt",
		},
	},
}

func assertFilename(t *testing.T, got *Ctx, expected string) {
	if got.Destination != expected {
		t.Errorf("expected %q, got %q", expected, got.Destination)
	}
}
func TestSetLocalFilename(t *testing.T) {
	c := &Ctx{Path: "output.txt"}
	c.SetLocalFilename()
	assertFilename(t, c, "output.txt")

	c = &Ctx{Path: "/foo/bar/output.txt"}
	c.SetLocalFilename()
	assertFilename(t, c, "output.txt")

	c = &Ctx{Path: "./foo/output.txt"}
	c.SetLocalFilename()
	assertFilename(t, c, "output.txt")

	c = &Ctx{Path: "output.txt", Filename: "foo"}
	c.SetLocalFilename()
	assertFilename(t, c, "foo")
}
