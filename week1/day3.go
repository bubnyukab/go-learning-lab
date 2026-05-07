package week1

import (
	"fmt"
)

func Day3() {
	fmt.Println("================== Day 3 =================")

	fmt.Println(Hello("world", ""))
}

const (
	spanish = "Spanish"
	french = "French"
	dutch = "Dutch"
	
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	dutchHelloPrefix = "Hallo, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case dutch:
		prefix = dutchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
