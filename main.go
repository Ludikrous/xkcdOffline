package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// define flags
	pictureFolder := flag.String("loc", "./xkcd_comics/", "Full path for the folder in which comics will be stored")

	// parse flags
	flag.Parse()

	// populate picture folder with all the xkcd comics
	getAllComics(pictureFolder)
}

func getAllComics(pictureFolder *string) {

}
