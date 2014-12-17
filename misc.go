package main

import (
	"fmt"
	"os"
	"strings"
)

// writes out a message and then exits with the status coded provided by
// status.  Since bail calls os.Exit, defered functions are not run.
func bail(status int, template string, args ...interface{}) {
	if !strings.HasSuffix(template, "\n") {
		template += "\n"
	}
	if status == 0 {
		fmt.Fprintf(os.Stdout, template, args...)
	} else {
		fmt.Fprintf(os.Stderr, template, args...)
	}
	os.Exit(status)
}
