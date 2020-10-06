/*
Name:    Matthew Stewardson
Date:    23-09-20
Version: 1.0.4
Desc:    Forth iteration of my link checker project. Bug fixes
Optional Features: Colour and Timeout
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/fatih/color"
	flag "github.com/spf13/pflag"
)

/*
Opens and reads the given file into a single string. This string is then
checked for url's via a regex pattern.
*/
func readFile(file string, choice bool) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(f)
	re := regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,4}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)
	result := re.FindAllString(text, -1)
	fmt.Println("File output")
	var temp = "test"
	var k = 0
	if choice == true {
		fmt.Print("[")
	}
	for i := range result {
		if temp == result[i] {
			//This is to ignore dupes
		} else {
			checkStatus(result[i], k, choice)
			k++
		}
		temp = result[i]
	}
	if choice == true {
		fmt.Println(" ]")
	}
}

//Does a GET request on the given string found within the file
func checkStatus(url string, i int, choice bool) {
	//Colours
	c := color.New(color.FgCyan)
	r := color.New(color.FgRed)
	g := color.New(color.FgGreen)
	i++
	//Timeout
	client := &http.Client{
		Timeout: 7 * time.Second,
	}
	response, err := client.Get(url)

	if choice == true {
		fmt.Print(" { url: '")
		if err != nil {
			r.Print(url)
			fmt.Print("': status '")
			r.Print("Error")
		} else {
			if response.StatusCode >= 200 && response.StatusCode <= 299 {
				g.Print(url)
				fmt.Print("': status '")
				g.Print(response.StatusCode)
			} else if response.StatusCode == 400 {
				r.Print(url)
				fmt.Print("': status '")
				r.Print(response.StatusCode)
			} else if response.StatusCode == 404 {
				r.Print(url)
				fmt.Print("': status '")
				r.Print(response.StatusCode)
			} else {
				c.Print(url)
				fmt.Print("': status '")
				c.Print(response.StatusCode)
			}
			defer response.Body.Close()
		}
		fmt.Print("' },")
	} else {
		if err != nil {
			r.Println(i, " -> [ERROR]   ", "URL: ", url)
		} else {
			if response.StatusCode >= 200 && response.StatusCode <= 299 {
				g.Println(i, " -> [GOOD]    ", response.StatusCode, "URL: ", url)
			} else if response.StatusCode == 400 {
				r.Println(i, " -> [BAD]     ", response.StatusCode, "URL: ", url)
			} else if response.StatusCode == 404 {
				r.Println(i, " -> [BAD]     ", response.StatusCode, "URL: ", url)
			} else {
				c.Println(i, " -> [UNKNOWN]  URL: ", url)
			}
			defer response.Body.Close()
		}
	}
}

//Setting up the optional version command
var version = flag.BoolP("version", "v", false, "prints out version info")

//JSON output command agr
var JSONout = flag.BoolP("json", "j", false, "prints link output in JSON format")

func main() {

	var JSONchoice = false

	flag.Parse()
	if *version == true {
		fmt.Println("LinkStatus version 0.1.5")
		return
	}
	if *JSONout == true {
		fmt.Println("JSON output selected")
		JSONchoice = true
	}
	if len(os.Args) == 1 {
		fmt.Println(`
Name: LinkStatus
Usage: go run LinkStatus.go filenames
Example: go run LinkStatus.go urls.txt
Version: go run LinkStatus.go -v or --version to check version.
JSON Format: go run LinkStatus.go -j or --json to output as JSON format
				   `)
		os.Exit(0)
	}
	fmt.Println("Checking files ", os.Args[1:])
	for _, file := range os.Args[1:] {
		readFile(file, JSONchoice)
	}
	os.Exit(0)
}
