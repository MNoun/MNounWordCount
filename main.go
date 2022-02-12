/************************
Created by: Mitchell Noun
Date created: 2/10/22
Class: COMP415 Emerging Languages
Assignment: Project 1
*************************/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Enter file name:") // ask user for file name
	var fileName string
	fmt.Scanln(&fileName)
	content, err := ioutil.ReadFile(fileName) // take in file contents
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reg, err := regexp.Compile("[^a-zA-Z0-9]+") // regular expression to exclude punctuation
	if err != nil {
		log.Fatal(err)
	}
	processedContent := reg.ReplaceAllString(string(content), " ") // replace all punctuation in content
	contentSlice := strings.Split(string(processedContent), " ")   // convert content to slices

	// function calls
	countWords(contentSlice)
	reportResults(countWords(contentSlice))
}

// counts the unique words in the file
func countWords(contentSlice []string) map[string]int {
	count := make(map[string]int) // map to count each word
	for _, v := range contentSlice {
		count[v]++ // iterate when a word is found
	}
	return count
}

// prints out the map
func reportResults(count map[string]int) {
	results := count
	fmt.Println("The results are: ")
	for word, count := range results {
		fmt.Println("Word:", word, "=>", "Count:", count)
	}
	return
}
