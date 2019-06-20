package main

import (
	"flag"
)

const URL string  = "https://xkcd.com/"
const JSON_PATH string = "/info.0.json"

func main() {
	// define flags
	pictureFolder := flag.String("loc", "./xkcd_comics/", "Full path for the folder in which comics will be stored")

	// parse flags
	flag.Parse()

	// populate picture folder with all the xkcd comics
	allComics := getAllComics(pictureFolder)
}

func getAllComics(pictureFolder *string) []comic {
	// get latest comic number
	currNumber := getHighestComicNum()

	// make a slice of comics and populate them
	allComics := make([]comic, currNumber)
	// iterate through all comics
	for num:=currNumber; num <= 0; num-- {
		allComics[num] = NewComic(num)
	}

	return allComics
}

func getHighestComicNum() int {
	return 2165		// TODO this method should be implemented later on
}
