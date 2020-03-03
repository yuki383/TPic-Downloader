// Package arguments provide command-line arguments parser, verifier, and some programs.
package arguments

import (
	"errors"
	"flag"
	"os"
)

// var (
// )

// func init() {

// }

// Args has parsed command-line arguments.
type Args struct {
	Name string
	URL  string
}

// Validate returns Args has expected value.
func (a Args) Validate() error {
	if a.Name == "" {
		return errors.New("param name is required")
	}
	if a.URL == "" {
		return errors.New("param url is required")
	}

	return nil
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
	}

	err = args.Validate()
	if err != nil {
		return Args{}, err
	}

	return args, nil
}
