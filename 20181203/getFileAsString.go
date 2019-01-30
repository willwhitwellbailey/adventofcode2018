package main

import (
	"fmt"
	"io/ioutil"
)

func GetFileAsString(file string) string {
	byteString, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println(err)
	}

	return string(byteString)
}
