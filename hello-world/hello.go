package main

import (
	"fmt"
)

const (
	spanish  = "Spanish"
	french   = "French"
	japanese = "Japanese"

	englishHelloPrefix  = "Hello, "
	spanishhHelloPrefix = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
	japaneseHelloPrefix = "今日, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishhHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
