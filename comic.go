package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
)

type comic struct {
	Number int

	Month int `json:"month,string"`
	Day   int `json:"day,string"`
	Year  int `json:"year,string"`

	Title     string `json:"title"`
	SafeTitle string `json:"safe_title"`
	Caption   string `json:"alt"`

	Address string `json:"img"`
}

// NewComic returns a comic object with all fields populated based on the given comic number.
// It fetches data for all these fields from the xkcd api (https://xkcd.com/json.html).
func NewComic(number int) (c comic) {
	// initialize comic struct
	c = comic{}
	c.Number = number

	// http GET request to xkcd api
	comicURL := fmt.Sprintf("http://www.xkcd.com/%d/info.0.json", number)
	resp, err := makeAPIRequest(comicURL)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal xkcd's response
	defer resp.Body.Close()
	err = jsoniter.NewDecoder(resp.Body).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// makeAPIRequest sends a get request to the given url and returns its response.
func makeAPIRequest(url string) (*http.Response, error) {
	c := http.Client{}

	// create an http request to send to xkcd
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "xkcdOffline")

	// send the request
	res, getErr := c.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	return res, nil
}
