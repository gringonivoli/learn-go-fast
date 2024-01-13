package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)
	// *p = 10
	// fmt.Printf("\nThe value p points to is: %v", *p)
	p = &i
	i = 777
	fmt.Printf("\nThe value p points to is: %v", *p)

	fmt.Println()

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice // pointer, linked
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	fmt.Println()

	functionThings()
}

func functionThings() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	// var result [5]float64 = square(thing1) // value, not reference, immutable
	var result [5]float64 = square(&thing1) // pointer, reference, mutable
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe thing1 is: %v", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	// func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing2)
	for index := range thing2 {
		thing2[index] = thing2[index] * thing2[index]
	}
	return *thing2
}
