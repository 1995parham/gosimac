package main

import (
	"fmt"

	"github.com/1995parham/gosimac/cmd"
)

// nolint: gocheckglobals
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Printf("gosimac %s, commit %s, built at %s\n", version, commit, date)

	cmd.Execute()
}
