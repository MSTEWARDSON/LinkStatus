
/*
Name:    Matthew Stewardson
Date:    22-09-20
Version: 1.0.0
Desc:    First iteration of my link checker project, so far you can give it a url
		 and it well return the status of that url.
Extra:   This is my first time using GO so I'm leaving comments to explain certain aspects of the code for reference.
*/
package main

import (
    "fmt"		//Standard formatted I/O functions
    "log"		//Simple logging package. (used for the log.fatal() function call)
    "net/http"	//This package provides http client and server implementations. (used to request the status on the url)
)

//Simple http status checker
func main() {

    response, err := http.Get("https://httpstat.us")									//Get req on the given url
    if err != nil {
        log.Fatal(err)																	//using log.Fatal to report on any incorrect url
    }

    fmt.Println("HTTP Response Status:", response.StatusCode, http.StatusText(response.StatusCode))

    if response.StatusCode >= 200 && response.StatusCode <= 299 {
        fmt.Println("HTTP Status is in the 200 - 299 range. It's good!")
	} else if response.StatusCode == 400{
		fmt.Println("HTTP Status is 400. It's not good!")
	} else if response.StatusCode == 404{
		fmt.Println("HTTP Status is 404. It's not good!")
	} else {
        fmt.Println("Unknown/Broken")
    }
}