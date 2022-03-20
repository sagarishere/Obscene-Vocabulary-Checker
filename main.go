package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// write your code here
	fileName := ""
	_, e := fmt.Scan(&fileName)
	if e != nil {
		return
	}
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	wholeFile := strings.ToLower(string(file))
	fmt.Println(wholeFile)
	// take a word as a user input
	// check if it is in the file
	sentence := ""
	prevSentence := ""
	// keep taking user input until user types "exit"
	for {
		_, er := fmt.Scan(&sentence)
		if er != nil {
			return
		}
		if sentence == "exit" {
			fmt.Println("Bye!")
			break
		}
		if sentence == prevSentence {
			continue
		}
		prevSentence = sentence
		// split the sentence and store it in a slice of words
		sliceOfWords := strings.Split(sentence, " ")
		for i, word := range sliceOfWords {
			if strings.Contains(wholeFile, strings.ToLower(word)) {
				sliceOfWords[i] = strings.Repeat("*", len(word))
			}
		}
		// join slice of words with spaces and print it
		fmt.Println(strings.Join(sliceOfWords, " "))
	}
}
