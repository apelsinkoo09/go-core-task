package main

import "fmt"

func Intersection(a, b []int) ([]int, bool) {
	intersectMap := make(map[int]bool)
	var intersectionSlice []int
	for _, i := range a {
		intersectMap[i] = true
	}
	for _, j := range b {
		if intersectMap[j] == true {
			intersectionSlice = append(intersectionSlice, j)
		}
	}
	if intersectionSlice == nil {
		return nil, false
	} else {
		return intersectionSlice, true
	}
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	result, flag := Intersection(a, b)
	if flag == true {
		fmt.Println(result)
	} else {
		fmt.Println("Intersection not found")
	}
}
