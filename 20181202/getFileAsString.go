package main

import (
	"fmt"
	"io/ioutil"
)

func GetFileAsString(file string) string {
	byteString, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	return string(byteString)
}
