package main

import (
	"testing"
)

//Tests the readFile functions ability to construct a []string
//when supplied a file
func TestReadFile(t *testing.T) {
	fakeResult := []string{"https://wiki.cdot.senecacollege.ca/w/api.php?action=rsd"}
	var result []string
	result = readFile("testReadFile.txt", false, 1, false, false)
	if result[0] != fakeResult[0] {
		t.Errorf("Result was incorrect")
	}
}

//Testing what happens when I give readFile an empty file
func TestEmptyFile(t *testing.T) {
	var result []string
	result = readFile("emptyFile.txt", false, 1, false, false)
	if result != nil {
		t.Errorf("Result was incorrect")
	}
}

//Test that will fail
func TestReadFileFail(t *testing.T) {
	fakeResult := []string{"https://wiki.cdot.senecacollege.ca/w/api.php?action=rsd"}
	var result []string
	result = readFile("testReadFile.txt", false, 1, false, false)
	if result[0] != fakeResult[0] {
		t.Errorf("Result was incorrect")
	}
}
