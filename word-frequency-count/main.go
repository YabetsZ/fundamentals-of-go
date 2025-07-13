package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Print("Enter the string to count: ")

	var input string
	// scanner because the input might be space separated text
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	
	fmt.Println(input)
	// Since the task prefers case-insensitive
	input = strings.ToLower(strings.TrimSpace(input))
	
	var count = make(map[string]int)
	
	for _, char := range input {
		if unicode.IsLetter(char) {
			count[string(char)] += 1
		}
	}
	fmt.Printf("%v", count)
}
