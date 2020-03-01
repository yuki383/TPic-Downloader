package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Args has parsed command-line arguments.
type Args struct {
	Name string
	URL  string
}

func main() {
	flag.Parse()
	args := Args{
		Name: flag.Arg(0),
		URL:  flag.Arg(1),
	}
	err := args.Validate()
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := http.Get(args.URL)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	t := time.Now().Local()
	b := MakeBasePath(args.Name, t)
	p := filepath.Join("data", b)
	abs, err := filepath.Abs(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create(abs)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	i, err := file.Write(body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("save successed bytes: %d\n", i)
}

// Validate returns Args has expected value.
func (a Args) Validate() error {
	if a.URL == "" {
		return errors.New("param url is required")
	}
	if a.Name == "" {
		return errors.New("param name is required")
	}

	return nil
}

// MakeBasePath return absolute base path using specified args.
// basePath format: <unix seconds>_<name>
func MakeBasePath(name string, date time.Time) string {
	min := date.Unix()
	path := fmt.Sprintf("%d_%s", min, name)

	return path
}
