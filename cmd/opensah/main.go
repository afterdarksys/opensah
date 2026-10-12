package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/afterdarksys/opensah/internal/suite"
)

const version = "0.1.0"

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	if len(args) == 0 {
		usage()
		return 2
	}
	switch args[0] {
	case "version", "--version":
		fmt.Println("opensah " + version)
		return 0
	case "help", "-h", "--help":
		usage()
		return 0
	case "tools":
		fmt.Print(suite.Lines())
		return 0
	case "ready":
		text, ok := suite.Ready(func(name string) error {
			_, err := exec.LookPath(name)
			return err
		})
		fmt.Print(text)
		if !ok {
			return 1
		}
		return 0
	default:
		usage()
		return 2
	}
}

func usage() {
	fmt.Fprint(os.Stderr, `opensah `+version+` — initial SAH tooling

  opensah version
  opensah tools
  opensah ready

tools prints the role of each SAH name. ready checks that sahlink, sahconnect,
and sahapi are on PATH. sahcore is unassigned and is not a failure. adssh is
reported and is not required for ready.
`)
}
