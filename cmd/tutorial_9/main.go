package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	// channels1()
	// channels2()
	moreRealisticChannel()
}

func moreRealisticChannel() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChikenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nText Sent: Found a deal on chicken at %v\n", website)
	case website := <-tofuChannel:
		fmt.Printf("\nEmail Sent: Found a deal on tofu at %v\n", website)
	}
}

func checkTofuPrices(aWebsite string, aChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_TOFU_PRICE {
			aChannel <- aWebsite
			break
		}
	}
}

func checkChikenPrices(aWebsite string, aChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			aChannel <- aWebsite
			break
		}
	}
}

func channels2() {
	var c = make(chan int, 5)
	go process(c)
	// fmt.Println(<-c) // waiting to the channel is set
	for value := range c {
		fmt.Println(value)
		time.Sleep(time.Second * 1)
	}
}

func channels1() {
	var c = make(chan int)
	go process(c)
	// fmt.Println(<-c) // waiting to the channel is set
	for value := range c {
		fmt.Println(value)
	}
}

func process(c chan int) {
	defer close(c)
	// c <- 123
	for i := 0; i < 5; i++ {
		c <- i
	}
	// close(c)
	fmt.Println("Exiting Process") // to see the diff between make(chan int) and make(chan int, 5)
}
