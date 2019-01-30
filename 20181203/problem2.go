package main

import (
	"fmt"
	"strings"
)

func Problem2() {
	fmt.Println("-----------------------------")

	inputs := strings.Split(GetFileAsString("input.txt"), "\n")

	fmt.Printf("last input claim number is %d\n", GetClaimNumber(inputs[len(inputs)-1]))

	// fmt.Println("Solution:")
	fmt.Println("-----------------------------")
}
