// Package arguments provide command-line arguments parser, verifier, and some programs.
package arguments

import (
	"errors"
	"flag"
	"os"
)

var (
	path = flag.String("p", "data", "install path of fetched image.")
	format = flag.String("f", "jpg", "specify image extension.\ndefault: jpg\naccept: jpg or png")
	size = flag.String("s", "large", "specify image size.\ndefault: large\naccept: small, midium, large, thumb")
)

// func init() {

// }

// Args has parsed command-line arguments.
type Args struct {
	Name string
	URL  string

	Path string
	Format string
	Size string
}

// Validate returns Args has expected value.
func (a Args) Validate() error {
	if a.Name == "" {
		return errors.New("param name is required")
	}
	if a.URL == "" {
		return errors.New("param url is required")
	}

	if !(a.Format == "jpg" || a.Format == "png") {
		return errors.New("invalid fomat")
	}

	if !sizeValidate(a.Size) {
		return errors.New("invalid size")
	}

	return nil
}

func sizeValidate(size string) bool {
	sizes := []string{ "small", "midium", "large", "thumb" }
	for _, s := range sizes {
		if size == s {
			return true
		}
	}

	return false
}

// Parser is
type Parser interface {
	Arg(i int) string
	Parse(arguments []string) error
}

var flagSet Parser = flag.CommandLine

// ParseFlags returns Args in command-line arguments.
func ParseFlags() (Args, error) {
	return parseFlags(flagSet)
}

func parseFlags(parser Parser) (Args, error) {
	err := parser.Parse(os.Args[1:])
	if err != nil {
		return Args{}, err
	}

	args := Args{
		Name: parser.Arg(0),
		URL:  parser.Arg(1),
		Path: *path,
		Format: *format,
		Size: *size,
	}

	err = args.Validate()
	if err != nil {
		return Args{}, err
	}

	return args, nil
}
