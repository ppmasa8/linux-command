package main

import (
	"io"
	"os"
)

func main() {
	rt := os.Args[1]
	cp := os.Args[2]

	src, err := os.Open(rt)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(cp)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
}
