package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	// ownerInfo owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
	// ownerInfo owner
}

func (self gasEngine) milesLeft() uint8 {
	return self.mpg * self.gallons
}

func (self electricEngine) milesLeft() uint8 {
	return self.kwh * self.mpkwh
}

type engine interface {
	milesLeft() uint8
}

type car struct {
	engine
}

// this function was lonly, so I create a car struct to stick it to :D
func (self car) canMakeIt(miles uint8) {
	if miles <= self.engine.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

// type owner struct {
// 	name string
// }

func main() {
	// var aGasEngine gasEngine = gasEngine{mpg: 22, gallons: 88, ownerInfo: owner{name: "Max"}}
	var aGasEngine gasEngine = gasEngine{mpg: 22, gallons: 88}
	var anElectricEngine electricEngine = electricEngine{kwh: 22, mpkwh: 88}
	aGasEngine.mpg = 11
	fmt.Printf("\nMPG: %v, Gallons: %v", aGasEngine.mpg, aGasEngine.gallons)
	fmt.Printf("\nTotal miles left in tank: %v", aGasEngine.milesLeft())
	fmt.Println()

	// is almost kind a nice...
	car{engine: aGasEngine}.canMakeIt(100)
	car{engine: anElectricEngine}.canMakeIt(200)

	// anonimous structure
	var anEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{33, 44}

	fmt.Printf("\nMPG: %v, Gallons: %v", anEngine2.mpg, anEngine2.gallons)
}
