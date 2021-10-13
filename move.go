package main

import (
	"log"
	"os"
)

func main() {
	rt := os.Args[1]
	mv := os.Args[2]

	if err := os.Rename(rt, mv); err != nil {
		log.Fatal(err)
	}
}
