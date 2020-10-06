/*
Name:    Matthew Stewardson
Date:    06-10-20
Version: 0.1.4
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

//Global Var's
var typeLink int = 1

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
	var k = 0
	for i := range result {
		if temp == result[i] {
			//This is to ignore dupes
		} else {
			checkStatus(result[i], k)
			k++
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

	switch typeLink {
	case 1:
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
	case 2:
		if err != nil {
			//Nothing
		} else {
			if response.StatusCode >= 200 && response.StatusCode <= 299 {
				g.Println(i, " -> [GOOD]    ", response.StatusCode, "URL: ", url)
			}
			defer response.Body.Close()
		}
	case 3:
		if err != nil {
			r.Println(i, " -> [ERROR]   ", "URL: ", url)
		} else {
			if response.StatusCode == 400 {
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

//All flag
var all = flag.BoolP("all", "a", false, "Prints out all types of responses")

//Good flag
var good = flag.BoolP("good", "g", false, "Prints out only good responses")

//Bad flag
var bad = flag.BoolP("bad", "b", false, "Prints out only bad responses")

func main() {
	flag.Parse()
	if *version == true {
		fmt.Println("LinkStatus version 0.1.4")
		return
	}
	if *all == true {
		fmt.Println("Outputting all types of links")
		typeLink = 1
	}
	if *good == true {
		fmt.Println("Outputting only good types of links")
		typeLink = 2
	}
	if *bad == true {
		fmt.Println("Outputting only bad types of links")
		typeLink = 3
	}
	if len(os.Args) == 1 {
		fmt.Println(`
Name: LinkStatus
Usage: go run LinkStatus.go filenames
Example: go run LinkStatus.go urls.txt
Version: go run LinkStatus.go -v or --version to check version.
All:  go run LinkStatus.go -a or --all to output all types of responses
Good: go run LinkStatus.go -g or --good to output only good types of responses
Bad:  go run LinkStatus.go -b or --bad to output only bad types of responses
				   `)
		os.Exit(-1)
	}
	fmt.Println("Checking files ", os.Args[1:])
	for _, file := range os.Args[1:] {
		readFile(file)
	}
}
