package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users []User `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Type string `xml:"type,attr"`
	Name string `xml:"name"`
	Social Social `xml:"social"`
}

type Social struct {
	XMLName 	xml.Name 	`xml:"social"`
	Facebook 	string 		`xml:"facebook"`
	Twitter 	string		`xml:"twitter"`
	Youtube		string		`xml:"youtube"`
}

func main() {
	// open the xml file
	xmlFile, err := os.Open("users.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened users.xml")

	defer xmlFile.Close()

	// Read our opened xmlFile as a byte array

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var users Users

	xml.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User: ", i+1)
		fmt.Println("User type: " + users.Users[i].Type)
		fmt.Println("User name: " + users.Users[i].Name)
		fmt.Println("Facebook: " + users.Users[i].Social.Facebook)
		fmt.Println("Twitter: " + users.Users[i].Social.Twitter)
		fmt.Println("Youtube: " + users.Users[i].Social.Youtube)
		fmt.Println()
	}
}