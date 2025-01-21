package main

import (
	"math"
	"regexp"
	"strconv"
)

func remainder(n int) bool {
	return n%2 == 0
}

func bitwise(n int) bool {
	return n&1 == 0
}

func round(n int) bool {
	return math.Abs(math.Round(float64(n)/2)*2-float64(n)) < 0.5
}

func cos(n int) bool {
	return math.Cos(float64(n)*math.Pi) > 0
}

func sin(n int) bool {
	return math.Abs(math.Sin(float64(n)*math.Pi/2)) < 0.5
}

func binStr(n int) bool {
	binaryStr := strconv.FormatInt(int64(n), 2)
	return binaryStr[len(binaryStr)-1] == '0'
}

func regex(n int) bool {
	match, _ := regexp.MatchString(`.*[02468]$`, strconv.Itoa(n))

	return match
}

func loop(n int) bool {
	even := true

	for i := 0; i < n; i++ {
		even = !even
	}

	return even
}

func recursive(n int) bool {
	if n == 0 {
		return true
	}

	if n == 1 {
		return false
	}

	return recursive(n - 2)
}

func str(n int) bool {
	numStr := strconv.Itoa(n)
	lastChar := numStr[len(numStr)-1]

	return lastChar == '0' || lastChar == '2' || lastChar == '4' || lastChar == '6' || lastChar == '8'
}
