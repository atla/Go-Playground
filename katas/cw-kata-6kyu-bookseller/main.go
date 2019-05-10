package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//...
}

// StockList function
func StockList(listArt []string, listCat []string) string {

	var out = ""

	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	for idx, element := range listCat {

		sum := 0

		for _, art := range listArt {
			if strings.HasPrefix(art, element) {
				count, err := strconv.Atoi(strings.Split(art, " ")[1])

				if err == nil {
					sum += count
				}
			}
		}
		if idx > 0 {
			out += " - "
		}
		out += fmt.Sprintf("(%s : %d)", element, sum)

	}
	return out
}
