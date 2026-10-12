package suite

import (
	"errors"
	"strings"
	"testing"
)

func TestCatalogKeepsSahcoreUnassigned(t *testing.T) {
	text := Lines()
	if !strings.Contains(text, "sahcore\tnot_built\tunassigned") {
		t.Fatalf("catalog:\n%s", text)
	}
	if !strings.Contains(text, "opensah\tbuilt\tinitial specific tooling") {
		t.Fatalf("opensah line missing:\n%s", text)
	}
	if strings.Contains(text, "renamed") {
		t.Fatal("catalog claims a rename")
	}
}

func TestReadyFailsOnlyMissingBuiltTools(t *testing.T) {
	look := func(name string) error {
		if name == "sahapi" || name == "adssh" {
			return errors.New("missing")
		}
		return nil
	}
	text, ok := Ready(look)
	if ok {
		t.Fatal("ready passed with sahapi missing")
	}
	if !strings.Contains(text, "sahapi\tbuilt\tmissing") {
		t.Fatalf("text:\n%s", text)
	}
	if !strings.Contains(text, "adssh\texternal\tmissing") {
		t.Fatalf("adssh line:\n%s", text)
	}
	if !strings.Contains(text, "sahcore\tnot_built\tno-binary") {
		t.Fatalf("sahcore line:\n%s", text)
	}
	_, ok = Ready(func(string) error { return nil })
	if !ok {
		t.Fatal("ready failed when built tools exist")
	}
}
