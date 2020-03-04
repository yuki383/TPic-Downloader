package main

import (
	"fmt"

	"tpic-fetcher/pkg/arguments"
	"tpic-fetcher/pkg/files"
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

	size, err := files.New("data", pics.Name, pics.Data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("save successed bytes: %d\n", size)
}
