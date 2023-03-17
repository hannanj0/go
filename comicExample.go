package main

import "fmt"

func main() {

	// example program that creates variables and assigns values to them using Go
	var publisher, writer, artist, title string
	var year, pageNumber int
	var grade float32

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "Dizzybooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist, "is published by", publisher, "In the year", year, ",it has a total of:", pageNumber, "pages and this copy has a grade of:", grade)

	fmt.Println("------------------------------------------------")

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	// publisher stays the same, so do not have to change. We will include to be diligent
	publisher = "Dizzybooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 9.0

	fmt.Println(title, "written by", writer, "drawn by", artist, "is published by", publisher, "In the year", year, ",it has a total of:", pageNumber, "pages and this copy has a grade of:", grade)

}
