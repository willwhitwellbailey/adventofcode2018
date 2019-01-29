package main

import (
	"fmt"
	"strconv"
	"strings"
)

func contains(slice []int64, n int64) bool {
	for _, num := range slice {
		if num == n {
			return true
		}
	}
	return false
}

func DuplicateFrequencyFinder() {
	fmt.Println("-----------------------------")
	inputs := strings.Split(GetFileAsString("input.txt"), "\n")

	const MAX_FREQUENCY_LOOP_COUNT int = 500

	var frequency int64 = 0
	var frequencies []int64                      // this is a slice
	frequencies = append(frequencies, frequency) // this includes the first frequency of 0
	foundDuplicate := false
	for loop := 0; !foundDuplicate && loop < MAX_FREQUENCY_LOOP_COUNT; loop++ {
		for i := 0; i < len(inputs); i++ {
			s := strings.Split(inputs[i], "\r")[0]

			// ParseInt automatically detects positive and negative signs - neat!
			number, x := strconv.ParseInt(s, 10, 0)
			if x != nil {
				fmt.Println("Parse error", x)
			}

			frequency += number

			if contains(frequencies, frequency) {
				foundDuplicate = true
				fmt.Println("Duplicate encountered", frequency)
				fmt.Println("Found within loop", loop)
				break
			}

			// Creating a binary tree here would be SUPER helpful
			// Or a map for O(1) searches
			// This is definitely the beginner brute-force method and
			//   my machine started slowing down around the 70th loop (took 144)
			frequencies = append(frequencies, frequency)
		}
	}

	fmt.Println(frequency)
	fmt.Println("-----------------------------")
}
