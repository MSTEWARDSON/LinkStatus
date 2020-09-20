package main

import (
	"fmt"		//implements formatted I/O
	"io/ioutil"	//I/O utility functions
)

func main() {
	data, err := ioutil.ReadFile("test.txt") //reads the file and returns a byte slice which is stored in data
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data)) //converts data to string and displays the content of the file
}