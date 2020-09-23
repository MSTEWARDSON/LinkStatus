
/*
Name:    Matthew Stewardson
Date:    23-09-20
Version: 1.0.1
Desc:    Second iteration of my link checker project. I added the ability to read from a given file, and check url's from the file
Extra:   This is my first time using GO so I'm leaving comments to explain certain aspects of the code for reference.
*/
package main

import (
    "fmt"		
    "log"		
	"net/http"	
	"io/ioutil"
	"regexp"
	flag "github.com/spf13/pflag"		//Using pflag here for added functionality compared to flag
	"os"
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

	for i := range(result) {
		checkStatus(result[i], i)
	}
}

//Does a GET request on the given string found within the file
func checkStatus(url string, i int) {
	i++
	response, err := http.Get(url)		

    if err != nil {
        log.Fatal(err)																	
    }
    if response.StatusCode >= 200 && response.StatusCode <= 299 {
		if i < 10 {
			fmt.Println(i, " -> [GOOD]    " , response.StatusCode, "URL: ", url)
		} else {
			fmt.Println(i, "-> [GOOD]    " , response.StatusCode, "URL: ", url)
		}
	} else if response.StatusCode == 400{
		if i < 10 {
			fmt.Println(i, " -> [BAD]     " , response.StatusCode, "URL: ", url)
		} else {
			fmt.Println(i, "-> [BAD]     " , response.StatusCode, "URL: ", url)
		}
	} else if response.StatusCode == 404{
		if i < 10 {
			fmt.Println(i, " -> [BAD]     " , response.StatusCode, "URL: ", url)
		} else {
			fmt.Println(i, "-> [BAD]     " , response.StatusCode, "URL: ", url)
		}
	} else {
		if i < 10 {
			fmt.Println(i, " -> [UNKNOWN] " , response.StatusCode, "URL: ", url)
		} else {
			fmt.Println(i, "-> [UNKNOWN] " , response.StatusCode, "URL: ", url)
		}
	}
}

//Setting up the optional version command
var version = flag.BoolP("version", "v", false, "prints out version info")

func main() {
	flag.Parse()

	if *version == true {
		fmt.Println("LinkStatus version 0.1.1")
		return
	}

	if len(os.Args) == 1 {
		fmt.Println(`
Name: LinkStatus
Usage: go run LinkStatus.go filenames
Example: go run LinkStatus.go urls.txt
Version: go run main.go -v or --version to check version.
				   `)
		os.Exit(-1)
	}

	fmt.Println("Checking files ", os.Args[1:])

	for _, file := range os.Args[1:] {
		readFile(file)
	}
}
