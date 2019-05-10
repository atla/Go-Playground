package main

import "strings"

func main() {}

// Accum ...
func Accum(s string) string {
	var result string
	for i, c := range s {
		result += strings.Title(strings.ToLower(strings.Repeat(string(c), i+1))) + "-"
	}
	return result[:len(result)-1]
}
