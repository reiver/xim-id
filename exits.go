package main

import (
	"fmt"
	"os"
)

func badrequest(format string, a ...interface{}) {
	var s string = fmt.Sprintf(format, a...)
	fmt.Fprintln(os.Stderr, s)

	const commandLineUsageError = 64
	os.Exit(commandLineUsageError)
}

func internalerror(format string, a ...interface{}) {
	var s string = fmt.Sprintf(format, a...)
	fmt.Fprintln(os.Stderr, s)

	const commandLineInternalSoftwareError = 70
	os.Exit(commandLineInternalSoftwareError)
}
