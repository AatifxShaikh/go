package main

import "fmt"

const hello = "Hello, "
const Spanish = "Spanish"
const helloForSpanish = "Hola, "
const French = "French"
const helloForFrench = "Bonjour, "
const Japanese = "Japanese"
const helloForJapanese = "Konnichiwa, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	// if language == Spanish {
	// 	return helloForSpanish + name
	// }
	// if language == French {
	// 	return helloForFrench + name
	// }
	//refactoring and adding switch
	prefix := hello
	switch language {
	case Spanish:
		prefix = helloForSpanish
	case French:
		prefix = helloForFrench
	case Japanese:
		prefix = helloForJapanese
	}
	return prefix + name
}
func main() {
	fmt.Println(Hello("world", "English"))
}
