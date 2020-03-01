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

func main() {
	args, err := parseFlags()
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

	time := time.Now()
	min := time.Local().Unix() / 60

	file, err := os.Create(fmt.Sprintf("data/%d_%s", min, args.Name))
	if err != nil {
		fmt.Println(err)
		fmt.Println("in os.Create()")
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

// Args has parsed command-line arguments.
type Args struct {
	Name string
	URL  string
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

func parseFlags() (Args, error) {
	flag.Parse()

	a := Args{
		Name: flag.Arg(0),
		URL:  flag.Arg(1),
	}

	err := a.Validate()
	if err != nil {
		return Args{}, err
	}

	return a, nil
}

// MakeFilePath return absolute path using specified args.
// baseURL format: <unix seconds>_<name>
func MakeFilePath(name string, date time.Time) (string, error) {
	min := date.Unix()
	rel := fmt.Sprintf("data/%d_%s", min, name)

	abs, err := filepath.Abs(rel)
	if err != nil {
		return "", err
	}

	return abs, nil
}
