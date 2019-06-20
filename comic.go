package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
	"time"
)


	//{
	//"month": "6",
	//"num": 2165,
	//"link": "",
	//"year": "2019",
	//"news": "",
	//"safe_title": "Millennials",
	//"transcript": "",
	//"alt": "Ironically, I've been having these same arguments for at least a decade now. I thought we would have moved on by now, but somehow the snide complaints about millennials continue.",
	//"img": "https://imgs.xkcd.com/comics/millennials.png",
	//"title": "Millennials",
	//"day": "19"
	//}



type comic struct {
	Number  int

	Month int `json:"month"`
	Day int `json:"day"`
	Year int `json:"year"`

	Title   string `json:"title"`
	SafeTitle string `json:"safe_title"`
	Caption string `json:"alt"`

	Address string `json:"img"`
}

func NewComic(number int) (c comic) {
	// initialize comic struct
	c = comic{}
	c.Number = number

	comicURL := fmt.Sprintf("http://www.xkcd.com/%d/info.0.json", number)
	req, err := makeAPIRequest(comicURL)

	return c
}

func makeAPIRequest(url string) (req, error) {

	c := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "xkcdOffline")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
}
