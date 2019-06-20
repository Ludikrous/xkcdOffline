package main

import (
	"fmt"
	"github.com/ludikrous/xkcdOffline/comic"
	_ "github.com/ludikrous/xkcdOffline/comic"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const URL string = "https://xkcd.com/"
const JSON_PATH string = "/info.0.json"

//func main() {
//	// define flags
//	pictureFolder := flag.String("loc", "./xkcd_comics/", "Full path for the folder in which comics will be stored")
//	// TODO flags to add:
//	// log into csv, high res, add title and captions to image, compile into pdf, start end date, start end number
//
//	// parse flags
//	flag.Parse()
//
//	// populate picture folder with all the xkcd comics
//	allComics := getAllComics(pictureFolder)
//
//	// download all comics into the given directory
//	for _, comic := range(allComics) {
//		download(comic, pictureFolder)
//	}
//}

func main() {
	// define flags
	pictureFolder := "/Users/dhanvee_ivaturi/Downloads/"

	// populate picture folder with all the xkcd comics
	allComics := getAllComics()
	// download all comics into the given directory
	for _, comic := range allComics {
		download(comic, &pictureFolder)
	}
}

func download(c comic.Comic, pictureFolder *string) {
	// get the image from xkcd servers
	response, e := http.Get(c.Address)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	//open a file for writing
	splitURL := strings.Split(c.Address, "/")
	fileEnding := splitURL[len(splitURL)-1]

	filename := fmt.Sprintf("%v/xkcdOffline/xkcd%d_%s", *pictureFolder, c.Number, fileEnding)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Saved comic #%d to disk.", c.Number)

}

func getAllComics() []comic.Comic {
	// get latest comic number
	currNumber := getHighestComicNum()

	// make a slice of comics and populate them
	allComics := make([]comic.Comic, currNumber)
	// iterate through all comics
	for num := currNumber; num > 0; num-- {
		c := comic.NewComic(num)
		allComics[num-1] = c
		fmt.Println(c)
	}

	return allComics
}

func getHighestComicNum() int {
	return 10 // TODO this method should be implemented later on
}
