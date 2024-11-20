package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	a := 42
	b := 052
	c := 0x2A
	d := 3.14
	e := "Golang"
	f := true
	g := 1 + 2i
	variables := []interface{}{a, b, c, d, e, f, g}
	var text string
	for _, varType := range variables {
		//fmt.Println("Type", reflect.TypeOf(varType))
		fmt.Printf("Type %T of %#v\n", varType, varType)
		symbol := fmt.Sprintf("%v", varType)
		text += symbol
	}
	fmt.Println(text)
	runeText := []rune(text)

	mid := len(runeText) / 2
	runeTextFinal := append(runeText[:mid])
	runeTextFinal = append([]rune("go_2204"))
	runeTextFinal = append(runeText[mid+1:])

	fmt.Println(runeText)
	finalText := string(runeTextFinal)

	hasher := sha256.New()
	hasher.Write([]byte(finalText))
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
}
