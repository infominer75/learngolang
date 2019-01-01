package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

/*
	A simple XKCD content fetcher.
	To get the JSON content for the current comic: https://xkcd.com/info.0.json
	To get the JSON content for comic 614: https://xkcd.com/614/info.0.json
	Sample response:
	{
		"month": "12",
		"num": 2085,
		"link": "",
		"year": "2018",
		"news": "",
		"safe_title": "arXiv",
		"transcript": "",
		"alt": "Both arXiv and archive.org are invaluable projects which, if they didn't exist, we would dismiss as obviously ridiculous and unworkable.",
		"img": "https://imgs.xkcd.com/comics/arxiv.png",
		"title": "arXiv",
		"day": "14"
	}
 */

 type ContentDescriptor struct {
 	Month int		`json:"month"`
 	Num int		`json:"num"`
 	Link string	`json:"link"`
 	Year int		`json:"year"`
 	News string	`json:"news"`
 	SafeTitle string	`json:"safe_title"`
 	Transcript string	`json:"transcript"`
 	AltText string		`json: "alt"`
 	ImageUrl string	`json:"img"`
 	Title string			`json:"title"`
 	Day string		`json:"day"`
 }

const templ  = `--------------------------content date : {{.Month}}/ {{.Day}}/{{.Year}}---------------------------
Image URL : {{.ImageUrl}}
Title: {{.Title}}
Transcript: {{.Transcript}}
Alt_Text: {{.AltText}}
---------------------------------------------------------------------------------------------------------------------------------------------
`

 func FetchContent(contentNumber string) (*ContentDescriptor, error) {
 		var comicUrl string
 		if strings.ToLower(contentNumber) == "current" {
			comicUrl = fmt.Sprintf("https://xkcd.com/info.0.json")
		} else {
			comicUrl = fmt.Sprintf("https://xkcd.com/%s/info.0.json", contentNumber)
		}
		resp, err := http.Get(comicUrl)

		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("Encountered HTTP status code %d for contentNumber %s", resp.StatusCode, contentNumber)
		}
		jsonDecoder := json.NewDecoder(resp.Body)
		var contentDescriptor ContentDescriptor
		jsonDecoder.Decode(&contentDescriptor)
		return &contentDescriptor, nil
 }

func main() {
	var contentDescriptor *ContentDescriptor
	var err error

	if len(os.Args )== 1 {
		contentDescriptor, err = FetchContent("current")
	} else {
		contentDescriptor, err = FetchContent(os.Args[1])
	}

	if err != nil {
		fmt.Printf("Encountered error when fetching content %v",err)
		os.Exit(-1)
	}
	fmt.Printf("Content descriptor : %v\n", *contentDescriptor)
	fmt.Println("Done")

	//print the JSON
	data, err := json.MarshalIndent(contentDescriptor, "", "\t")
	if err != nil {
		fmt.Printf("Error when marshalling data to json : %v", err)
		os.Exit(-2)
	}
	fmt.Printf("%s\n", data)
	//attempt to print using a report template
	var report = template.Must(template.New("report").Parse(templ))
	if err := report.Execute(os.Stdout, contentDescriptor); err != nil {
		fmt.Println("Failed to parse %v content to report", *contentDescriptor)
		os.Exit(-3)
	}


}