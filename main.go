package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"tpic-fetcher/pkg/arguments"
)

func main() {
	args, err := arguments.ParseFlags()
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := http.Get(args.URL)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
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

// MakeBasePath return absolute base path using specified args.
// basePath format: <unix seconds>_<name>
func MakeBasePath(name string, date time.Time) string {
	min := date.Unix()
	path := fmt.Sprintf("%d_%s", min, name)

	return path
}
