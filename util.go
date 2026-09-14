package main

import (
	"fmt"
	"io"

	"github.com/cli/go-gh/v2"
)

// runCli runs the GitHub CLI, capturing stdOut and stdErr
func runCli(wo, we io.Writer, args ...string) error {
	stdOut, stdErr, err := gh.Exec(args...)
	if wo != nil {
		fmt.Fprint(wo, stdOut.String())
	}
	if we != nil {
		fmt.Fprint(we, stdErr.String())
	}
	return err
}
