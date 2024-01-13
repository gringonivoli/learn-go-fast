package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func main() {
	// someGenericsWithSlices()
	someGenericsWithChimi()
}

func someGenericsWithChimi() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("\n%+v", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("\n%+v", purchases)
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("\nError: %v\n", err.Error())
	}
	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}

func someGenericsWithSlices() {
	var intSlice = []int{1, 2, 3}
	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumIntSlice(intSlice))
	fmt.Println(sumSlice[int](intSlice))
	fmt.Println(sumFloat32Slice(float32Slice))
	fmt.Println(sumSlice[float32](float32Slice))

	fmt.Println(isEmpty(float32Slice))
	fmt.Println(isEmpty(intSlice))
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func sumSlice[T int | float32](slice []T) T {
	var sum T
	for _, value := range slice {
		sum += value
	}
	return sum
}

func sumIntSlice(slice []int) int {
	var sum int
	for _, value := range slice {
		sum += value
	}
	return sum
}

func sumFloat32Slice(slice []float32) float32 {
	var sum float32
	for _, value := range slice {
		sum += value
	}
	return sum
}
