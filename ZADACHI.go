package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func getNumberOfChannels() (int, error) {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return 0, fmt.Errorf("Invalud input, try again")
	}
	switch {
	case num < 0:
		return 0, fmt.Errorf("Negative number, try again")
	case num > 10:
		return 0, fmt.Errorf("Too much value, try again")
	}

	return num, nil
}

func randValueGen(numOfChannel int, ch chan int) {
	val := rand.Intn(10)
	ch <- val
	close(ch)
}

func merge(channels []chan int) chan int {
	var wg sync.WaitGroup
	mergedChannel := make(chan int)

	out := func(i <-chan int) {
		for j := range i {
			mergedChannel <- j
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, b := range channels {
		go out(b)
	}
	go func() {
		wg.Wait()
		close(mergedChannel)
	}()
	return mergedChannel
}

func main() {
	var numOfChannels int
	var err error

	for {
		fmt.Println("Enter number of channels to 10:")
		numOfChannels, err = getNumberOfChannels()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		break
	}
	fmt.Println("Number of channels = ", numOfChannels)

	chs := make([]chan int, numOfChannels)

	for i := 0; i < numOfChannels; i++ {
		chs[i] = make(chan int)
		go randValueGen(i, chs[i])
	}

	result := merge(chs)
	fmt.Print("Data from channels: ")
	for num := range result {
		fmt.Print(num)
	}
}
