package main

import (
	"fmt"
	"image"
)

type comic struct {
	number  int
	address string
	title   string
	image   image.Image
}

func New(number int) (c comic) {
	// initialize comic struct
	c = comic{}
	c.number = number
	c.address = fmt.Sprintf("http://www.xkcd.com/%d", number)

	// get data from xkcd servers
	c.populateData()

	// return populated struct
	return c
}

func (c comic) populateData() {

}
