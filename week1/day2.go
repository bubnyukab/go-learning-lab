package week1

import (
	"fmt"
	"strings"
	"golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		if _, ok := m[word]; !ok {
			m[word] = 1
		} else {
			m[word] += 1
		}
	}
	return m
}

func Day2() {
	// maps exercise
	wc.Test(wordCount)

	// function closures exercise
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}

func fibonacci() func() int {
	n1 := 0
	n2 := 1
	return func() int {
		result := n1
		n1, n2 = n2, n1+n2
		return result
	}
}

