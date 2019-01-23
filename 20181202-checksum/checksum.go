package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getFileAsString(file string) string {
	byteString, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	return string(byteString)
}

func createMap(s string) map[string]int {
	// create map of character + count of character within string
	// ie ccabcacc -->
	// {
	//   a: 2,
	//   b: 1,
	//   c: 5,
	// }

	m := make(map[string]int)

	for _, c := range s {
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

func main() {
	inputs := strings.Split(getFileAsString("input.txt"), "\n")

	var m map[string]int
	var doublesCount, triplesCount int = 0, 0
	for i := 0; i < len(inputs); i++ {
		fmt.Printf("input %d : %s\n", i, inputs[i])
		m = createMap(inputs[0])
		fmt.Printf("map %d with length %d : %v\n", i, len(m), m)

		if evalCharCount(m, 2) {
			doublesCount++
		}

		if evalCharCount(m, 3) {
			triplesCount++
		}
	}

	// fmt.Printf("doubles: %d triples %d\n", doublesCount, triplesCount)
	fmt.Println("checksum is ", doublesCount*triplesCount)
}
