package main

import (
	"fmt"
	"os"
	"path/filepath"

	"tpic-fetcher/pkg/arguments"
	"tpic-fetcher/pkg/pictures"
)

func main() {
	args, err := arguments.ParseFlags()
	if err != nil {
		fmt.Println(err)
		return
	}

	pics, err := pictures.New(args.Name, args.URL)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := filepath.Join("data", pics.Name)
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

	i, err := file.Write(pics.Data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("save successed bytes: %d\n", i)
}
