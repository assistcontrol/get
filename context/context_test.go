package context

import (
	"os"
	"testing"
)

func compareCtx(t *testing.T, got, expected *Ctx) {
	if got == nil {
		t.Fatal("expected context, got nil")
	}
	if got.Path != expected.Path {
		t.Errorf("expected Path to be %q, got %q", expected.Path, got.Path)
	}
	if got.Force != expected.Force {
		t.Errorf("expected Force to be %v, got %v", expected.Force, got.Force)
	}
	if got.Save != expected.Save {
		t.Errorf("expected Save to be %v, got %v", expected.Save, got.Save)
	}
	if got.Filename != expected.Filename {
		t.Errorf("expected Filename to be %q, got %q", expected.Filename, got.Filename)
	}
}

func TestNew(t *testing.T) {
	os.Args = []string{"get", "-o", "output.txt", "https://example.com"}
	expected := &Ctx{
		Path:     "https://example.com",
		Force:    false,
		Save:     true,
		Filename: "output.txt",
	}

	c, err := New()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	compareCtx(t, c, expected)
}
