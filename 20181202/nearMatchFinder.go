package main

import (
	"fmt"
	"strings"
)

// if we've gotten here, we already know that the strings are the same length
func similarChars(s1 string, s2 string) string {
	// would be better if a Builder was used, but let's keep it
	var simChars strings.Builder
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			fmt.Fprint(&simChars, string(s1[i]))
		}
	}
	return simChars.String()
}

func main() {
	inputs := strings.Split(GetFileAsString("input.txt"), "\n")

	for i := 0; i < len(inputs); i++ {
		for j := i + 1; j < len(inputs); j++ {
			// check for length mismatch
			if len(inputs[i]) > len(inputs[j]) || len(inputs[i]) < len(inputs[j]) {
				continue
			}

			diffCount := 0
			for k := 0; k < len(inputs[i]); k++ {
				if inputs[i][k] != inputs[j][k] {
					diffCount++
				}

				if diffCount > 2 {
					break
				}
			}

			if diffCount == 1 {
				fmt.Printf("Found almost matches: %d - %s and %d - %s\n", i, inputs[i], j, inputs[j])
				fmt.Println("Similar characters are ", similarChars(inputs[i], inputs[j]))
			}
		}
	}
}
