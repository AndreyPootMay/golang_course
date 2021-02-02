package main

import (
	"encoding/xml"
	"io/ioutil"
)

type Notes struct {
	To		string `xml:"to"`
	From	string `xml:"from"`
	Heading	string `xml:"heading"`
	Body	string `xml:"body"`
}

func main() {
	var notes [2]Notes

	// Creating two notes
	notes[0].To = "Apollo"
	notes[0].From = "Rocky"
	notes[0].Heading = "Fight"
	notes[0].Body = "The fight is 5pm"

	notes[1].To = "Niger"
	notes[1].From = "Rocky"
	notes[1].Heading = "Fight"
	notes[1].Body = "The fight is 7pm"

	// Creating an file (in the function only go prefix and indent)
	file, _ := xml.MarshalIndent(notes, " ", " ")

	// Writing to the file
	_ = ioutil.WriteFile("notes.xml", file, 0644)
}
