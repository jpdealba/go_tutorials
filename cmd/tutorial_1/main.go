package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	fmt.Println(intNum)

	var floatNum float32 = 3.141592
	fmt.Println(floatNum)

	var str string = "Hello" + " World!"
	newStr := "Hello"

	fmt.Println(str)
	fmt.Println(newStr)

	fmt.Println(len("Hello World!"))
	fmt.Println(utf8.RuneCountInString("Hello World!"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBool bool = true
	fmt.Println(myBool)

	// this is only usefull when dealing when obvious datatypes, but when dealing with func responses or that
	// are not obvious, it is better to use the var keyword and add the datatype
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// this is a constant, and it cannot be changed, it has to be initialized
	const myConst int = 10
	fmt.Println(myConst)

}
