package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

/*
Opens and reads the given file into a single string. This string is then
checked for url's via a regex pattern.
*/
func readFile(file string, jsonC bool, linkC int) {
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
	if jsonC == true {
		fmt.Print("[")
	}
	for i := range result {
		if temp == result[i] {
			//This is to ignore dupes
		} else {
			checkStatus(result[i], k, jsonC, linkC)
			k++
		}
		temp = result[i]
	}
	if jsonC == true {
		fmt.Println(" ]")
	}
}
