package main

import (
	"fmt"
	"strings"
)

func compareClaims(lhs map[string]bool, rhs map[string]bool) bool {
	// true if they overlap even just 1 inch

	for k := range lhs {
		_, present := rhs[k]
		if present {
			return true
		}
	}

	return false
}

func isSingleton(claim map[string]bool, existingClaimMaps map[int64]map[string]bool) (bool, map[int64]bool) {
	isSingleton := true
	overlays := make(map[int64]bool)
	for k := range existingClaimMaps {
		if compareClaims(claim, existingClaimMaps[k]) {
			// this means there is an overlap with an existing claim
			// can't just stop on the first overlay because there may be others

			// boolean becomes true
			isSingleton = false
			overlays[int64(k)] = true
		}
	}

	return isSingleton, overlays
}

func removeOverlays(singletons map[int64]bool, overlays map[int64]bool) map[int64]bool {
	for key := range overlays {
		if _, present := singletons[key]; present {
			delete(singletons, key)
		}
	}

	return singletons
}

func FindSingletons() {
	fmt.Println("-----------------------------")

	inputs := strings.Split(GetFileAsString("input.txt"), "\n")

	existingClaimMaps := make(map[int64]map[string]bool) // map of claim number -> maps with keys set to GetMapKey
	singletons := make(map[int64]bool)
	for i := 0; i < len(inputs); i++ {
		claim := GetClaimMap(inputs[i])
		isSingleton, overlays := isSingleton(claim, existingClaimMaps)
		if isSingleton {
			singletons[GetClaimNumber(inputs[i])] = true
		} else {
			removeOverlays(singletons, overlays)
		}
		existingClaimMaps[GetClaimNumber(inputs[i])] = claim
	}

	fmt.Println("Solution:")
	for k, _ := range singletons {
		fmt.Printf("  singleton claim %v\n", k)
	}
	fmt.Println("-----------------------------")
}
