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

func TestCatalogIncludesEarlyTools(t *testing.T) {
	text := Lines()
	if !strings.Contains(text, "sahsync\toptional\tfile copy and sync through the systemapi agent") {
		t.Fatalf("sahsync line missing:\n%s", text)
	}
	if !strings.Contains(text, "sahxfer\tearly\tbulk one-file transfer on its own udp data plane, authorized by a grant") {
		t.Fatalf("sahxfer line missing:\n%s", text)
	}
	if !strings.Contains(text, "sah2grpc\tearly\tgrpc debug bridge for sah decisions") {
		t.Fatalf("sah2grpc line missing:\n%s", text)
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

func TestReadyIgnoresMissingEarlyTools(t *testing.T) {
	look := func(name string) error {
		if name == "sahsync" || name == "sahxfer" {
			return errors.New("missing")
		}
		return nil
	}
	text, ok := Ready(look)
	if !ok {
		t.Fatalf("ready failed with only early tools missing:\n%s", text)
	}
	if !strings.Contains(text, "sahsync\toptional\tmissing") {
		t.Fatalf("sahsync line:\n%s", text)
	}
	if !strings.Contains(text, "sahxfer\tearly\tmissing") {
		t.Fatalf("sahxfer line:\n%s", text)
	}
	if !strings.Contains(text, "sah2grpc\tearly\tno-binary") {
		t.Fatalf("sah2grpc line:\n%s", text)
	}
}
