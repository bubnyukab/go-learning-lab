package week1

import (
	"fmt"
	"strings"
)

func Day4() {
	fmt.Println("================== Day 4 =================")

	result := Repeat("yourmom", 69)
	fmt.Println(result)
}

func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}
