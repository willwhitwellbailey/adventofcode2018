package main

import (
	"strconv"
	"strings"
)

func GetClaimNumber(claim string) int64 {
	number := strings.Split(claim, " ")[0]
	n, _ := strconv.ParseInt(number[1:], 10, 0)
	return n
}

func GetClaimCornerCoordinates(claim string) (int64, int64) {
	claimCorner := strings.Split(claim, " ")[2]
	claimCorner = claimCorner[:len(claimCorner)-1] // removes colon
	claimCorners := strings.Split(claimCorner, ",")
	cornerX, _ := strconv.ParseInt(claimCorners[0], 10, 0)
	cornerY, _ := strconv.ParseInt(claimCorners[1], 10, 0)

	return cornerX, cornerY
}

func GetClaimDims(claim string) (int64, int64) {
	claimDims := strings.Split(strings.Split(claim, " ")[3], "x")
	width, _ := strconv.ParseInt(claimDims[0], 10, 0)
	height, _ := strconv.ParseInt(strings.Split(claimDims[1], "\r")[0], 10, 0) // BAH!

	return width, height
}

func GetMapKeyInt64(i int64, j int64) string {
	return strconv.FormatInt(i, 10) + "," + strconv.FormatInt(j, 10)
}

func GetMapKeyInt(i int, j int) string {
	return GetMapKeyInt64(int64(i), int64(j))
}
