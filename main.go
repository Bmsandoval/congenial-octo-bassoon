package main

import (
	"github.com/bmsandoval/congenial-octo-bassoon/internal"
	"github.com/bmsandoval/congenial-octo-bassoon/internal/modules"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("no file name provided. please provide a file to read")
		os.Exit(1)
	}

	userAccessCount := modules.UserAccessCount{}
	userAccessCount.Setup()

	uploadsOverFifty := modules.UploadsOverFifty{}
	uploadsOverFifty.Setup()

	internal.ParseFile(os.Args[1], &userAccessCount, &uploadsOverFifty)
}

