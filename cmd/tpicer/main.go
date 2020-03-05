package main

import (
	"fmt"
	"os"

	"tpic-fetcher/pkg/arguments"
	"tpic-fetcher/pkg/files"
	"tpic-fetcher/pkg/pictures"
)

func main() {
	args, err := arguments.ParseFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pics, err := pictures.New(args.Name, args.URL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	size, err := files.New("data", pics.Name, pics.Data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("save successed bytes: %d\n", size)
	os.Exit(0)
}
