package main

import "encoding/xml"

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id string `xml:"id,attr"`
	Content string  `xml:"content"`
	Author string `xml:"author"`
	Xml string `xml:",innerxml"`
}

type Author struct {
	Id string `xml:"id",attr`
	Name string `xml:",chardata"`
}

func main() {
	xmlFile, err := 
}