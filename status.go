package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
)

//Does a GET request on the given string found within the file
func checkStatus(url string, i int, jsonC bool, linkC int) {
	//Colours
	c := color.New(color.FgCyan)
	r := color.New(color.FgRed)
	g := color.New(color.FgGreen)
	i++
	//Timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := client.Get(url)

	if jsonC == true {
		fmt.Print(" { url: '")
		if err != nil {
			r.Print(url)
			fmt.Print("': status '")
			r.Print("Error")
		} else {
			if response.StatusCode >= 200 && response.StatusCode <= 299 {
				g.Print(url)
				fmt.Print("': status '")
				g.Print(response.StatusCode)
			} else if response.StatusCode == 400 {
				r.Print(url)
				fmt.Print("': status '")
				r.Print(response.StatusCode)
			} else if response.StatusCode == 404 {
				r.Print(url)
				fmt.Print("': status '")
				r.Print(response.StatusCode)
			} else {
				c.Print(url)
				fmt.Print("': status '")
				c.Print(response.StatusCode)
			}
			defer response.Body.Close()
		}
		fmt.Print("' },")
	} else {
		switch linkC {
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
}
