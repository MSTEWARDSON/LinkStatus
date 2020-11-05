/*
Name:    Matthew Stewardson
Date:    09-10-20
Version: 0.1.6
Desc:    Cleaning up the code/fixing bugs
Optional Features: Colour and Timeout
*/
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
Opens and reads the given file into a single string. This string is then
checked for urls via a regex pattern.
*/
func readFile(file string, jsonC bool, typeC int, ignore bool, telescope bool) []string {
	mainFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(mainFile)
	re := regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,4}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)
	if telescope == true {
		re = regexp.MustCompile(`https?:\/\/localhost:[0-9]{1,5}\/([-a-zA-Z0-9()@:%_\+.~#?&\/=]*)`)
	}
	result := re.FindAllString(text, -1)

	if ignore {
		var tempList []string
		ignoreList := ignoreURL("ignore.txt")
		fmt.Println("ignored urls ", ignoreList)
		fmt.Println("----------------------------------------------------------------------------------")
		for _, link := range result {
			valid := true
			for _, url := range ignoreList {
				if strings.HasPrefix(link, url) {
					valid = false
					break
				}
			}
			if valid {
				tempList = append(tempList, link)
			}
		}
		result = tempList
	}
	return result
}

/*
Reads the ignore test file and builds a string of certain urls to ignore
from the main file given.
*/
func ignoreURL(fileName string) []string {
	var ignoreList []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("^(#|https?://)")
	for scanner.Scan() {
		if !re.Match(scanner.Bytes()) {
			fmt.Println("Ignore file invalid")
			os.Exit(1)
		}
		if line := scanner.Text(); string(line[0]) != "#" {
			ignoreList = append(ignoreList, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ignoreList
}
