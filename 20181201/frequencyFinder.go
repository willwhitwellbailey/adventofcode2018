package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FrequencyFinder() {
	fmt.Println("-----------------------------")
	inputs := strings.Split(GetFileAsString("input.txt"), "\n")

	var frequency int64 = 0
	for i := 0; i < len(inputs); i++ {
		s := strings.Split(inputs[i], "\r")[0]

		// ParseInt automatically detects positive and negative signs - neat!
		number, x := strconv.ParseInt(s, 10, 0)
		if x != nil {
			fmt.Println("Parse error", x)
		}

		frequency += number
	}

	fmt.Println("Frequency Finder", frequency)
	fmt.Println("-----------------------------")
}
