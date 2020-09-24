/*
Name:    Matthew Stewardson
Date:    23-09-20
Version: 1.0.3
Desc:    Third iteration of my link checker project. I added colour and optional version commands
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
func readFile(file string) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	text := string(f)
	re := regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,4}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)
	result := re.FindAllString(text, -1)
	fmt.Println("File output")
	var temp = "test"
	for i := range result {
		if temp == result[i] {

		} else {
			checkStatus(result[i], i)
		}
		temp = result[i]
	}
}

//Does a GET request on the given string found within the file
func checkStatus(url string, i int) {

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

//Setting up the optional version command
var version = flag.BoolP("version", "v", false, "prints out version info")

func main() {
	flag.Parse()

	if *version == true {
		fmt.Println("LinkStatus version 0.1.2")
		return
	}

	if len(os.Args) == 1 {
		fmt.Println(`
Name: LinkStatus
Usage: go run LinkStatus.go filenames
Example: go run LinkStatus.go urls.txt
Version: go run LinkStatus.go -v or --version to check version.
				   `)
		os.Exit(-1)
	}

	fmt.Println("Checking files ", os.Args[1:])

	for _, file := range os.Args[1:] {
		readFile(file)
	}
}
