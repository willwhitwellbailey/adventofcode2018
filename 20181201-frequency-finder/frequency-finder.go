package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getFileAsString(file string) string {
	byteString, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	return string(byteString)
}

func main() {
	inputs := strings.Split(getFileAsString("input.txt"), "\n")

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

	fmt.Println(frequency)
}
