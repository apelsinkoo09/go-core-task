package main

import (
	"fmt"
	"math/rand"
)

func sliceExample(slice []int) []int {
	var evenSlice []int
	for _, i := range slice {
		if i%2 == 0 {
			evenSlice = append(evenSlice, i)
		}
	}
	return evenSlice
}

func addElements(slice []int, a int) []int {
	var newSlice []int
	newSlice = append(slice, a)
	return newSlice
}

func copySlice(slice []int) []int {
	copiedSlice := slice
	return copiedSlice
}

func removeElement(slice []int, a int) ([]int, error) {
	element := a
	if element >= len(slice) {
		return nil, fmt.Errorf("Invalid input")
	} else {
		newElement := append(slice[:element], slice[element+1:]...)
		return newElement, nil
	}
}

func main() {
	var originalSlice []int
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		originalSlice = append(originalSlice, a)
	}
	fmt.Println(originalSlice)
	evenSlice := sliceExample(originalSlice)
	fmt.Println(evenSlice)
	newSlice := addElements(originalSlice, 5)
	fmt.Println(newSlice)
	copiedSlice := copySlice(originalSlice)
	fmt.Println(copiedSlice)
	deletedElementSLice, err := removeElement(originalSlice, 3)
	if err != nil {
		fmt.Errorf("Failed to delete item:", err)
	} else {
		fmt.Println(deletedElementSLice)
	}
}
