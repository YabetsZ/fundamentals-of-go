package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func ChangeToCorrectForm(s string) string {
	var result strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) {
			result.WriteRune(unicode.ToLower(r))
		}
	}
	return result.String()
}

func main() {
	fmt.Print("Give the string to check: ")
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
		input = strings.TrimSpace(input)
		if input != "" {
			break
		}
		fmt.Println("Inalid input: The input shouldn't be an empty string.")
		fmt.Print("Give the string to check, again: ")
	}

	input = ChangeToCorrectForm(input)

	leftPtr, rightPtr := 0, len(input)-1
	for leftPtr < rightPtr {
		if input[leftPtr] != input[rightPtr] {
			fmt.Println("\n✖️ The given string is *NOT* a palindrome")
			return
		}
		leftPtr++
		rightPtr--
	}
	fmt.Println("\n✔️ The given string is a palindrome")
}
