package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if err := os.Mkdir(flag.Arg(0), 0777); err != nil {
		log.Fatal(err)
	}
}
