package main

import "fmt"

func main() {
	fmt.Printf("Hello world")

	var intNum int = 14
	fmt.Println(intNum)

	var floatNum float32 = 6
	fmt.Println(floatNum)

	var myString string = "abce"
	fmt.Println(myString)
	arrays()
	chars()
}
func arrays() {
	var intArr [3]int32
	intArr[1] = 23
	fmt.Println(intArr[0])
}

func chars() {
	myString := "resume"
	var indexed = myString[0]
	fmt.Printf("%v ,%T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
}

//structs define your own type
//go routines allows to run multiple function concurrently
