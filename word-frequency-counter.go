/*
Task:  Word Frequency Count
Write a Go function that takes a string as input and returns a dictionary containing the frequency of each word in the string.
Treat words in a case-insensitive manner and ignore punctuation marks.
[Optional]: Write a test for your function

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var word_count = make(map[string]int) // Map

	fmt.Printf("Enter words: ")

	// Get input from users.
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	words := input.Text()

	// fmt.Printf("%q", strings.Fields(words)) // This does not cater for the punctuation marks.

	// Using regular expression to get only the words
	re := regexp.MustCompile(`\w+`)
	// re := regexp.MustCompile(`[A-Za-zs1-9]+`) // Another variation to what is above.
	words_arr := re.FindAllString(words, -1)

	// Count all the words. All words are converted to lowercase to make it case-insensitive.
	for _, val := range words_arr {
		word_count[strings.ToLower(val)] += 1
	}

	fmt.Println(word_count)
}
