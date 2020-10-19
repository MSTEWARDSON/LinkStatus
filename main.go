/*
Name:    Matthew Stewardson
Date:    09-10-20
Version: 0.1.6
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
	var version = flag.BoolP("version", "v", false, "Prints out version info")
	var JSONout = flag.BoolP("json", "j", false, "Prints link output in JSON format")
	var all = flag.BoolP("all", "a", false, "Prints out all types of responses")
	var good = flag.BoolP("good", "g", false, "Prints out only good responses")
	var bad = flag.BoolP("bad", "b", false, "Prints out only bad responses")
	var ignore = flag.BoolP("ignore", "i", false, "Ignore certain url patterns")
	var JSONchoice = false
	var typeLink int = 1
	var result []string
	var temp = "test"
	var k = 0
	var total = 0

	flag.Parse()
	if *version == true {
		fmt.Println("LinkStatus version 0.1.6")
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
			result = readFile(file, JSONchoice, typeLink, *ignore)
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
	os.Exit(0)
}
