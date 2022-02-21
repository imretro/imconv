package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var help bool

func main() {
	flag.Parse()
	if help {
		printHelp()
		return
	}

	source := flag.Arg(0)
	dest := flag.Arg(1)

	if source == "" {
		exitOnError(errors.New("source not defined"))
	} else if dest == "" {
		exitOnError(errors.New("dest not defined"))
	}
}

func init() {
	flag.BoolVar(&help, "help", false, "print this message and exit")
}

func printHelp() {
	fmt.Fprintf(os.Stderr, "usage: imretro-converter [FLAGS] <source> <dest>\n\nFLAGS:\n")
	flag.PrintDefaults()
}

func exitOnError(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %v\n\n", err)
	printHelp()
	os.Exit(1)
}
