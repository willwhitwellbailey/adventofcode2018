package main

import (
	"fmt"
	"strings"
)

func addClaim(claim string, cloth map[string]int, maxW int, maxH int) (map[string]int, int, int) {
	cornerX, cornerY := GetClaimCornerCoordinates(claim)
	dimW, dimH := GetClaimDims(claim)

	if cornerX+dimW > int64(maxW) {
		maxW = int(cornerX + dimW)
	}

	if cornerY+dimH > int64(maxH) {
		maxH = int(cornerY + dimH)
	}

	for i := cornerX; i < cornerX+dimW; i++ {
		for j := cornerY; j < cornerY+dimH; j++ {
			cloth[GetMapKeyInt64(i, j)]++
		}
	}

	return cloth, maxW, maxH
}

func establishClaims(claims []string) (map[string]int, int, int) {
	// go through each input to increment "cut"s
	cloth := make(map[string]int)
	var maxW, maxH int
	for i := 0; i < len(claims); i++ {
		cloth, maxW, maxH = addClaim(claims[i], cloth, maxW, maxH)
	}

	return cloth, maxW, maxH
}

func getStats(cloth map[string]int, maxW int, maxH int) (int, int) {
	multipleClaims := 0
	zeroClaims := 0
	for i := 0; i < maxW; i++ {
		for j := 0; j < maxH; j++ {
			mapKey := GetMapKeyInt(i, j)
			inchClaimCount := cloth[mapKey]
			if inchClaimCount > 1 {
				multipleClaims++
			}

			if inchClaimCount == 0 {
				zeroClaims++
			}
		}
	}

	return multipleClaims, zeroClaims
}

func FindOverlaps() {
	fmt.Println("-----------------------------")

	inputs := strings.Split(GetFileAsString("input.txt"), "\n")

	cloth, maxW, maxH := establishClaims(inputs)

	fmt.Printf("The cloth has width %d and height %d\n", maxW, maxH)
	// go through each inch to increment if used multiple times

	multipleClaims, zeroClaims := getStats(cloth, maxW, maxH)

	fmt.Printf("Solution: %d inches have multiple claims\n", multipleClaims)
	fmt.Printf("Extra   : %d inches have zero claims\n", zeroClaims)
	fmt.Println("-----------------------------")
}
