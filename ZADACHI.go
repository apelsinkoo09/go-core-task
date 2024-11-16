package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func (s *StringIntMap) Add(key string, value int) {
	s.data[key] = value
}
func (s *StringIntMap) Remove(key string) {
	delete(s.data, key)
}
func (s *StringIntMap) Copy() map[string]int {
	copy := make(map[string]int)
	for key, value := range s.data {
		copy[key] = value
	}
	return copy
}
func (s *StringIntMap) Exists(key string) bool {
	_, exist := s.data[key]
	return exist
}
func (s *StringIntMap) Get(key string) (int, bool) {
	value, exist := s.data[key]
	if exist {
		return value, true
	} else {
		return 0, false
	}
}
func main() {
	karta := StringIntMap{
		data: make(map[string]int),
	}
	karta.Add("Formula", 1)
	karta.Add("Boeing", 747)
	karta.Add("VAZ", 2101)
	karta.Add("T", 90)
	fmt.Println(karta.data)

	karta.Remove("VAZ")
	fmt.Println(karta.data)

	copiedKarta := karta.Copy()
	fmt.Println("This is copy of original karta:", copiedKarta)

	existedValue := karta.Exists("Boeing")
	if existedValue == true {
		fmt.Println("Boeing 747 is existed")
	} else {
		fmt.Println("Boeing RIP")
	}

	geittingValue, exist := karta.Get("Formula")
	if exist {
		fmt.Println("Formula -", geittingValue)
	} else {
		fmt.Println("Enything")
	}
}
