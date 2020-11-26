/*
Name:    Matthew Stewardson
Date:    24-11-20
Version: 0.1.7
Desc:    Adding testing framework
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	flag "github.com/spf13/pflag"
)

func main() {
	var version = flag.BoolP("version", "v", false, "Prints out version info")
	var JSONout = flag.BoolP("json", "j", false, "Prints link output in JSON format")
	var all = flag.BoolP("all", "a", false, "Prints out all types of responses")
	var good = flag.BoolP("good", "g", false, "Prints out only good responses")
	var bad = flag.BoolP("bad", "b", false, "Prints out only bad responses")
	var ignore = flag.BoolP("ignore", "i", false, "Ignore certain url patterns")
	var telescope = flag.BoolP("telescope", "t", false, "Read data from telescope local host")
	var JSONchoice = false
	var typeLink int = 1
	var result []string
	var temp = "test"
	var k = 0
	var total = 0
	var tele = false

	flag.Parse()
	if *version == true {
		fmt.Println("LinkStatus version 0.1.7")
		return
	}
	if *JSONout == true {
		fmt.Println("JSON output selected")
		JSONchoice = true
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
	if *telescope == true {
		fmt.Println("Reading telescope post output")
		tele = true
	}
	if len(os.Args) == 1 {
		fmt.Println(`
==================================================================================
||                                LINK STATUS                                   ||
==================================================================================
Default: ./LinkStatus [filenames]
Example: ./LinkStatus urls.txt
Version: ./LinkStatus -v or --version to check version.
JSON: 	 ./LinkStatus -j or --json to output as JSON format
All:  	 ./LinkStatus -a or --all to output all types of responses
Good: 	 ./LinkStatus -g or --good to output only good types of responses
Bad:  	 ./LinkStatus -b or --bad to output only bad types of responses
Ignore:  ./LinkStatus test.txt -i or --ignore to ignore certain url patterns
==================================================================================
				   `)
		os.Exit(0)
	}
	fmt.Println("==================================================================================")
	fmt.Println("Command List: ", os.Args[1:])
	fmt.Println("----------------------------------------------------------------------------------")
	for _, file := range os.Args[1:] {
		if file[0] != '-' {
			if tele == true {
				telescopeParse(file)
				file = "tData.txt"
			}
			result = readFile(file, JSONchoice, typeLink, *ignore, tele)
			if JSONchoice == true {
				fmt.Print("[")
			}
			for i := range result {
				if temp == result[i] {
					//This is to ignore dupes
				} else {
					checkStatus(result[i], k, JSONchoice, typeLink)
					k++
					total++
				}
				temp = result[i]
			}
			if JSONchoice == true {
				fmt.Println(" ]")
			}
		}
	}
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Println("Total URL Count: ", total)
	fmt.Println("==================================================================================")
	//Remove the temp files
	err := os.Remove("telescopeData.json")
	if err != nil {
		log.Fatal(err)
	}
	e := os.Remove("tData.txt")
	if e != nil {
		log.Fatal(e)
	}
	os.Exit(0)
}

//Telescope struct which lays out the json data for easy storage
type Telescope struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func telescopeParse(file string) {
	fmt.Println("Telescope Parsing")
	fmt.Println(file)
	fmt.Println("----------------------------------------------------------------------------------")

	//Get the json data in a file
	out, _ := os.OpenFile("telescopeData.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer out.Close()
	resp, err := http.Get(file)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println(err)
	}

	jsonFile, err := os.Open("telescopeData.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	var teleData []Telescope

	err = json.Unmarshal(byteValue, &teleData)
	if err != nil {
		fmt.Println(err)
	}

	telescopeFile, err := os.OpenFile("tData.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(teleData); i++ {
		telescopeFile.Write([]byte("http://localhost:3000" + teleData[i].URL + " "))
	}

	telescopeFile.Close()
	jsonFile.Close()
}
