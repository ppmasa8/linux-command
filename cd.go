package main

import (
	"log"
	"os"
)

func main() {
	chDir := os.Args[1]
	//prevDir, _ := filepath.Abs(".")
	if err := os.Chdir(chDir); err != nil {
		log.Fatal(err)
	}
}
