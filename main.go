/*
Name:    Matthew Stewardson
Date:    09-10-20
Version: 0.1.5
Desc:    Cleaning up the code/fixing bugs
Optional Features: Colour and Timeout
*/
package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

func main() {
	var version = flag.BoolP("version", "v", false, "prints out version info")
	var JSONout = flag.BoolP("json", "j", false, "prints link output in JSON format")
	var all = flag.BoolP("all", "a", false, "Prints out all types of responses")
	var good = flag.BoolP("good", "g", false, "Prints out only good responses")
	var bad = flag.BoolP("bad", "b", false, "Prints out only bad responses")
	var JSONchoice = false
	var typeLink int = 1

	flag.Parse()
	if *version == true {
		fmt.Println("LinkStatus version 0.1.5")
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
	if len(os.Args) == 1 {
		fmt.Println(`
Name: LinkStatus
Usage: go run LinkStatus.go filenames
Example: go run LinkStatus.go urls.txt
Version: go run LinkStatus.go -v or --version to check version.
JSON Format: go run LinkStatus.go -j or --json to output as JSON format
All:  go run LinkStatus.go -a or --all to output all types of responses
Good: go run LinkStatus.go -g or --good to output only good types of responses
Bad:  go run LinkStatus.go -b or --bad to output only bad types of responses
				   `)
		os.Exit(0)
	}

	fmt.Println("Checking files ", os.Args[1:])
	for _, file := range os.Args[1:] {
		if file[0] != '-' {
			readFile(file, JSONchoice, typeLink)
		}
	}
	os.Exit(0)
}
