package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if err := os.Remove(flag.Arg(0)); err != nil {
		log.Fatal(err)
	}
}
