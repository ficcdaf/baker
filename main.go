package main

import (
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		// if errors.Is(err, fs.ErrExist) {
		// 	return
		// }
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	fmt.Println(len(os.Args), " arguments.")
	l := len(os.Args)
	if l != 2 {
		fmt.Println("Usage: baker [one file || dir]")
		os.Exit(1)
	}
	srcP := os.Args[1]
	i, err := os.Stat(srcP)
	check(err)
	if i.IsDir() {
		// We handle a directory differently
		dstP := fmt.Sprintf("%s.d.bak", srcP)
		err := os.CopyFS(dstP, os.DirFS(srcP))
		check(err)
	} else {
		src, err := os.Open(srcP)
		check(err)
		defer src.Close()
		dstP := fmt.Sprintf("%s.bak", srcP)
		// create file if doesn't exist
		dst, err := os.Create(dstP)
		check(err)

		_, err = io.Copy(dst, src)
		check(err)

		err = dst.Sync()
		check(err)
	}
}
