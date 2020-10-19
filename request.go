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
	"net/http"
	"time"

	"github.com/fatih/color"
)

//Does a GET request on the given string found within the file
func checkStatus(url string, i int, jsonC bool, typeC int) {
	//Colours
	cyan := color.New(color.FgCyan)
	red := color.New(color.FgRed)
	green := color.New(color.FgGreen)
	i++
	//Timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	//Request
	response, err := client.Get(url)

	if jsonC == true {
		fmt.Print(" { url: '")
		if err != nil {
			red.Print(url)
			fmt.Print("': status '")
			red.Print("Error")
		} else {
			if response.StatusCode >= 200 && response.StatusCode <= 299 {
				green.Print(url)
				fmt.Print("': status '")
				green.Print(response.StatusCode)
			} else if response.StatusCode == 400 {
				red.Print(url)
				fmt.Print("': status '")
				red.Print(response.StatusCode)
			} else if response.StatusCode == 404 {
				red.Print(url)
				fmt.Print("': status '")
				red.Print(response.StatusCode)
			} else {
				cyan.Print(url)
				fmt.Print("': status '")
				cyan.Print(response.StatusCode)
			}
			defer response.Body.Close()
		}
		fmt.Print("' },")
	} else {
		switch typeC {
		case 1:
			if err != nil {
				red.Println(i, "[ERROR]   ", "URL: ", url)
			} else {
				if response.StatusCode >= 200 && response.StatusCode <= 299 {
					green.Println(i, "[GOOD]    ", response.StatusCode, "URL: ", url)
				} else if response.StatusCode == 400 {
					red.Println(i, "[BAD]     ", response.StatusCode, "URL: ", url)
				} else if response.StatusCode == 404 {
					red.Println(i, "[BAD]     ", response.StatusCode, "URL: ", url)
				} else {
					cyan.Println(i, "[UNKNOWN]  URL: ", url)
				}
				defer response.Body.Close()
			}
		case 2:
			if err != nil {
				//Nothing
			} else {
				if response.StatusCode >= 200 && response.StatusCode <= 299 {
					green.Println(i, "[GOOD]    ", response.StatusCode, "URL: ", url)
				}
				defer response.Body.Close()
			}
		case 3:
			if err != nil {
				red.Println(i, "[ERROR]   ", "URL: ", url)
			} else {
				if response.StatusCode == 400 {
					red.Println(i, "[BAD]     ", response.StatusCode, "URL: ", url)
				} else if response.StatusCode == 404 {
					red.Println(i, "[BAD]     ", response.StatusCode, "URL: ", url)
				} else {
					if response.StatusCode <= 200 && response.StatusCode >= 299 {
						cyan.Println(i, "[UNKNOWN]  URL: ", url)
					}
				}
				defer response.Body.Close()
			}
		}
	}
}
