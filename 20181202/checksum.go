package main

import (
	"fmt"
	"strings"
)

func createMap(str string) map[string]int {
	// create map of character + count of character within string
	// ie ccabcacc -->
	// {
	//   a: 2,
	//   b: 1,
	//   c: 5,
	// }

	m := make(map[string]int)

	for _, c := range str {
		m[string(c)]++
	}

	return m
}

func evalCharCount(m map[string]int, count int) bool {
	for _, v := range m {
		if v == count {
			return true
		}
	}
	return false
}

func Checksum() {
	fmt.Println("-----------------------------")
	inputs := strings.Split(GetFileAsString("input.txt"), "\n")

	var doublesCount, triplesCount int = 0, 0
	for i := 0; i < len(inputs); i++ {
		m := createMap(inputs[i])

		if evalCharCount(m, 2) {
			doublesCount++
		}

		if evalCharCount(m, 3) {
			triplesCount++
		}
	}

	fmt.Printf("doubles: %d triples %d\n", doublesCount, triplesCount)
	fmt.Println("checksum is ", doublesCount*triplesCount)
	fmt.Println("-----------------------------")
}
