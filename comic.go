package main

import (
	"fmt"
)

type comic struct {
	Number  int
	Title   string
	Caption string
	Address string
}

func New(number int) (c comic) {
	// initialize comic struct
	c = comic{}
	c.Number = number

	comicURL := fmt.Sprintf("http://www.xkcd.com/%s", c.Title)

	return c
}

func (c comic) populateData() {

}
