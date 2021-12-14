package main

import "fmt"

const englishPrefix = "Hello "

func Hello(name string) string {
	if name == "" {
		return "Hello World"
	}
	return fmt.Sprintf("%s%s", englishPrefix, name)
}

func Bye() string {
	return "Bye"
}

func main() {
	fmt.Println(Hello("Pierrot"))
}
