/*
Task: Palindrome Check
Write a Go function that takes a string as input and checks whether it is a palindrome or not.
A palindrome is a word, phrase, number, or other sequence of characters that reads
the same forward and backward (ignoring spaces, punctuation, and capitalization).
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

func ReverseString(str string) string {
	/*
		This reverses the order of a string.
	*/

	output := ""
	for _, val := range str {
		output = string(val) + output
	}
	return output
}

func main() {
	fmt.Printf("Enter text: ")

	// Get user input.
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	words := input.Text()

	// Find words and make them case-insensitive while neglecting punctuations.
	re := regexp.MustCompile(`\w+`)
	words_arr := re.FindAllString(strings.ToLower(words), -1)

	// Join all words together.
	words_str := strings.Join(words_arr, "")

	fmt.Println("Is this a palindrome?", ReverseString(words_str) == words_str)
}
