package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 3

func main() {
	// basicGolang()
	var chickenChannel = make(chan string)
	var websites []string = []string{"walmart.com", "cotsco.com", "wholefoods.com"}
	for _, website := range websites {
		go checkChickenPrices(website, chickenChannel)
	}
	sendMessage(chickenChannel)
	close(chickenChannel)

}

func checkChickenPrices(website string, chickenChannel chan<- string) {
	var tryCount int32 = 0
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice float32 = rand.Float32()*20 + 1
		if chickenPrice <= MAX_CHICKEN_PRICE {
			fmt.Printf("Tried %d times\n", tryCount)
			chickenChannel <- website
			break
		} else {
			tryCount++
			fmt.Println(tryCount)
		}
	}
}

func sendMessage(chickenChannel <-chan string) {
	fmt.Printf("Found a deal on chicken at %s\n", <-chickenChannel)
}
