package main

import "strings"

func main() {}

func bandNameGenerator(word string) string {

	if word[0] == word[len(word)-1] {
		return strings.Title(word + word[1:len(word)])
	}
	return "The " + strings.Title(word)
}
