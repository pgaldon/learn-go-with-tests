package main

import "fmt"

const englishPrefix = "Hello "
const spanishPrefix = "Hola "

func Hello(name string, lang string) string {
	var prefix string
	if lang == "" {
		lang = "EN"
	}
	if lang == "EN" {
		prefix = englishPrefix
	} else if lang == "SP" {
		prefix = spanishPrefix
	}
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s%s", prefix, name)
}

func Bye() string {
	return "Bye"
}

func main() {
	fmt.Println(Hello("Pierrot", "EN"))
}
