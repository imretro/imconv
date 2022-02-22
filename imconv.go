package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	imretro "github.com/imretro/go"
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

	sourceFile, err := os.Open(source)
	if err != nil {
		exitOnError(err)
	}
	defer sourceFile.Close()

	sourceImage, _, err := image.Decode(sourceFile)
	if err != nil {
		exitOnError(err)
	}

	destFile, err := os.Create(dest)
	if err != nil {
		exitOnError(err)
	}
	defer destFile.Close()

	switch ext := strings.ToLower(filepath.Ext(dest)); ext {
	case ".gif":
		err = gif.Encode(destFile, sourceImage, nil)
	case ".imretro":
		err = imretro.Encode(destFile, sourceImage, imretro.EightBit)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(destFile, sourceImage, nil)
	case ".png":
		err = png.Encode(destFile, sourceImage)
	default:
		err = fmt.Errorf("Unsupported image format: %v", ext)
	}
	if err != nil {
		exitOnError(err)
	}
}

func init() {
	flag.BoolVar(&help, "help", false, "print this message and exit")
}

func printHelp() {
	fmt.Fprintf(os.Stderr, "usage: imconv [FLAGS] <source> <dest>\n\nFLAGS:\n")
	flag.PrintDefaults()
}

func exitOnError(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %v\n\n", err)
	printHelp()
	os.Exit(1)
}
