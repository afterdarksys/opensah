package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestVersionAndTools(t *testing.T) {
	if run([]string{"version"}) != 0 {
		t.Fatal("version")
	}
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	code := run([]string{"tools"})
	_ = w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	if code != 0 || !strings.Contains(buf.String(), "sahcore\tnot_built\tunassigned") {
		t.Fatalf("code %d\n%s", code, buf.String())
	}
	if !strings.Contains(buf.String(), "sahsync\toptional\tfile copy and sync through the systemapi agent") {
		t.Fatalf("sahsync line missing:\n%s", buf.String())
	}
	if run([]string{"nope"}) != 2 {
		t.Fatal("usage")
	}
}
