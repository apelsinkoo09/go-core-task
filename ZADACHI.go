package main

import (
	"fmt"
	"math/rand"
)

func writeToChan(ch1 chan uint8) {
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		numU8 := uint8(num)
		ch1 <- numU8
	}
	close(ch1)
}

func convertToFloat64InCube(ch1 chan uint8, ch2 chan float64) {
	for num := range ch1 {
		numF := float64(num)
		numCube := numF * numF * numF
		ch2 <- numCube
	}
	close(ch2)
}

func main() {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	go writeToChan(ch1)
	go convertToFloat64InCube(ch1, ch2)

	for result := range ch2 {
		fmt.Println(result)
	}
}
