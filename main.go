package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
)

var force bool

func check(err error) {
	if err != nil {
		if errors.Is(err, fs.ErrExist) && !force {
			fmt.Printf("Target exists and -f flag not set!\n")
			os.Exit(1)
		}
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func checkExists(dstP string) bool {
	_, err := os.Stat(dstP)
	if !errors.Is(err, os.ErrNotExist) {
		if !force {
			fmt.Printf("Target exists and -f flag not set!\n")
			os.Exit(1)
		} else {
			return true
		}
	}
	return false
}

// TODO: Add a -n flag which creates a NEW backup file instead of overwriting
func main() {
	forcePtr := flag.Bool("f", false, "Force; overwrite the output file if it exists.")
	flag.Parse()
	force = *forcePtr
	args := flag.Args()
	l := len(args)
	if l != 1 {
		fmt.Println("Usage: baker [file || dir]")
		os.Exit(1)
	}
	srcP := args[0]
	i, err := os.Stat(srcP)
	check(err)
	if i.IsDir() {
		// We handle a directory differently
		dstP := fmt.Sprintf("%s.d.bak", srcP)
		if checkExists(dstP) {
			err := os.RemoveAll(dstP)
			check(err)
		}
		err = os.CopyFS(dstP, os.DirFS(srcP))
		check(err)
	} else {
		dstP := fmt.Sprintf("%s.bak", srcP)
		checkExists(dstP)
		src, err := os.Open(srcP)
		check(err)
		defer src.Close()
		// create file if doesn't exist
		dst, err := os.Create(dstP)
		check(err)

		_, err = io.Copy(dst, src)
		check(err)

		err = dst.Sync()
		check(err)
	}
}
