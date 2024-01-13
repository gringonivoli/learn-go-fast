package main

import (
	"fmt"
	"strings"
)

func main() {
	// var myString = ("résumé")
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for index, value := range myString {
		fmt.Println(index, value)
	}
	fmt.Printf("\nThe length of 'myString' is %v", len(myString))

	var aRune = 'a'
	fmt.Printf("\naRune = %v", aRune)

	// strings are inmutable

	var strSlice = []string{"m", "a", "x", "i"}
	var strBuilder strings.Builder
	for index := range strSlice {
		strBuilder.WriteString(strSlice[index]) // faster that asign a new string every time...
	}
	fmt.Printf("\n%v", strBuilder.String())
}
