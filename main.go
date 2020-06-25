package main

import (
	"github.com/bmsandoval/congenial-octo-bassoon/internal"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("no file name provided. please provide a file to read")
		os.Exit(1)
	}

	internal.ParseFile(os.Args[1])
}

