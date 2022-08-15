package main

import (
	"fmt"
	"os"

	"github.com/w-haibara/cuc/pkg/cmd/root"
)

type exitCode int

const (
	exitOK    exitCode = 0
	exitError exitCode = 1
)

func main() {
	code := mainRun()
	os.Exit(int(code))
}

func mainRun() exitCode {
	rootCmd := root.NewCmdRoot()
	if cmd, err := rootCmd.ExecuteC(); err != nil {
		fmt.Fprintln(os.Stderr, "failed:", cmd.Name(), "\n", err.Error())
		return exitError
	}

	return exitOK
}
