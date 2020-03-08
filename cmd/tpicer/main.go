package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"tpic-fetcher/pkg/arguments"
	"tpic-fetcher/pkg/files"
	"tpic-fetcher/pkg/pictures"
)

func setURL(url string, format string, size string) string {
	u := strings.SplitAfter(url, "?")[0]

	return fmt.Sprintf("%sformat=%s&name=%s", u, format, size)
}

func setName(name string, extend string) (string, error) {
	ex, err := regexp.Compile(`\.[jpg|png]$`)
	if err != nil {
		return "", err
	}

	if ex.Match([]byte(name)) {
		splited := strings.Split(name, ".")
		exclude := splited[:len(splited)-1]

		name = strings.Join(exclude, "")
	}

	return name + "." + extend, nil

}

func do() error {
	args, err := arguments.ParseFlags()
	if err != nil {
		return err
	}

	url := setURL(args.URL, args.Format, args.Size)

	pics, err := pictures.New(args.Name, url)
	if err != nil {
		return err
	}

	name, err := setName(pics.Name, args.Format)
	if err != nil {
		return err
	}

	s, err := files.New(args.Path, name, pics.Data)
	if err != nil {
		return err
	}

	fmt.Printf("write success! size: %v", s)
	return nil
}

func main() {
	err := do()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
