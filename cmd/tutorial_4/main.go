package main

import (
	"fmt"
	"time"
)

func main() {
	doArrayThings()
	separator()
	doSliceThings()
	separator()
	doMapThings()
	separator()
	performanceTest()
}

func performanceTest() {
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

func separator() {
	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Println()
}

func doArrayThings() {
	var intArr1 [3]int32
	intArr1[1] = 123
	fmt.Println(intArr1[0])
	fmt.Println(intArr1[1:3])

	// memory location...
	fmt.Println(&intArr1[0])
	fmt.Println(&intArr1[1])

	// others initializations...
	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr2[1])

	intArr3 := [3]int32{1, 2, 3}
	fmt.Println(intArr3[2])

	intArr4 := [...]int32{1, 2, 3}
	fmt.Println(intArr4[2])

	for index, value := range intArr1 {
		fmt.Printf("Index: %v, Value: %v", index, value)
	}
}

func doSliceThings() {
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("The length is %v with capacity %v", len(intSlice3), cap(intSlice3))
}

func doMapThings() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["???"]) // returns the default value if key not exists, that will be 0 in this case
	var age, ok = myMap2["XXX"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("Value not found...")
	}

	for name, age := range myMap2 { // without order
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}
}
