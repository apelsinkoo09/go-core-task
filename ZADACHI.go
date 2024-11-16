package main

import "fmt"

func Duplicate(s1 []string, s2 []string) []string {
	dup := make(map[string]bool)
	var s3 []string
	for _, i := range s2 {
		dup[i] = true
	}
	for _, j := range s1 {
		if dup[j] != true {
			s3 = append(s3, j)
		}
	}
	return s3
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	slice3 := Duplicate(slice1, slice2)
	fmt.Println(slice3)
}
