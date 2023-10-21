package main

import (
	"fmt"
	"io"
	"os"
)

type CLI struct {
	outStream io.Writer
	errStream io.Writer
}

const (
	ExitSuccess = iota
	ExitError
)

func main() {
	cli := &CLI{
		outStream: os.Stdout,
		errStream: os.Stderr,
	}

	os.Exit(cli.Run(os.Args))
}

func (cli *CLI) Run(args []string) int {

	tag := ""
	if len(args) > 1 {
		tag = args[1]
	}

	next, err := GetNextVersion(tag)
	if err != nil {
		errmsg := fmt.Sprintf("Getting release tag error: %s.", err.Error())
		fmt.Fprintln(cli.errStream, errmsg)
		return ExitError
	}
	fmt.Fprintln(cli.outStream, next)

	return ExitSuccess
}
