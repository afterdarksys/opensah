// Package suite is the first OpenSAH tool: a local index of the SAH
// groundwork. It does not provision machines or open a shell.
package suite

import (
	"fmt"
	"strings"
)

// Tool is one named part of the suite.
type Tool struct {
	Name  string
	Role  string
	State string
	Bin   string
}

// Catalog is the role map. sahcore has no binary because its role is unassigned.
func Catalog() []Tool {
	return []Tool{
		{Name: "opensah", Role: "initial specific tooling", State: "built"},
		{Name: "sahlink", Role: "stdio and opt-in socket capture for a program this tool starts", State: "built", Bin: "sahlink"},
		{Name: "sahconnect", Role: "ssh and systemapi client", State: "built", Bin: "sahconnect"},
		{Name: "sahapi", Role: "local api service", State: "built", Bin: "sahapi"},
		{Name: "sahsync", Role: "file copy and sync through the systemapi agent", State: "optional", Bin: "sahsync"},
		{Name: "sahxfer", Role: "bulk one-file transfer on its own udp data plane, authorized by a grant", State: "early", Bin: "sahxfer"},
		{Name: "sah2grpc", Role: "grpc debug bridge for sah decisions", State: "early"}, // no binary yet
		{Name: "sahcore", Role: "unassigned", State: "not_built"},
		{Name: "adssh", Role: "execution shell", State: "external", Bin: "adssh"},
		{Name: "systemapi", Role: "control plane", State: "external"},
		{Name: "computeapi", Role: "compute catalog", State: "external"},
	}
}

// Lines renders the catalog for humans and scripts.
func Lines() string {
	var b strings.Builder
	for _, tool := range Catalog() {
		fmt.Fprintf(&b, "%s\t%s\t%s\n", tool.Name, tool.State, tool.Role)
	}
	return b.String()
}

// Ready reports whether the built tools are on PATH. A missing external
// binary is reported and does not fail the check. sahcore is not a failure.
func Ready(look func(string) error) (string, bool) {
	var b strings.Builder
	ok := true
	for _, tool := range Catalog() {
		if tool.Bin == "" {
			fmt.Fprintf(&b, "%s\t%s\tno-binary\n", tool.Name, tool.State)
			continue
		}
		if err := look(tool.Bin); err != nil {
			fmt.Fprintf(&b, "%s\t%s\tmissing\n", tool.Name, tool.State)
			if tool.State == "built" {
				ok = false
			}
			continue
		}
		fmt.Fprintf(&b, "%s\t%s\tok\n", tool.Name, tool.State)
	}
	return b.String(), ok
}
