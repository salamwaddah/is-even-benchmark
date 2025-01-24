package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	functions := map[string]func(int) bool{
		"remainder": remainder,
		"bitwise":   bitwise,
		"round":     round,
		"cos":       cos,
		"sin":       sin,
		"binStr":    binStr,
		"regex":     regex,
		"loop":      loop,
		"recursive": recursive,
		"str":       str,
	}

	data := map[string][]time.Duration{}

	for name, fn := range functions {
		data[name] = benchmark(name, fn)
	}

	filename := "results_100_000.txt"
	_ = saveToTxt(filename, data)
}

func benchmark(name string, fn func(int) bool) []time.Duration {
	runs := 100
	results := make([]time.Duration, runs)

	for r := 0; r < runs; r++ {
		start := time.Now()
		for i := 0; i < 100_000; i++ {
			_ = fn(i)
		}
		results[r] = time.Since(start)
	}

	avgTime := time.Duration(0)
	for _, t := range results {
		avgTime += t
	}
	avgTime /= time.Duration(runs)
	fmt.Printf("%s took an average of %s\n", name, avgTime)

	return results
}

func saveToTxt(filename string, data map[string][]time.Duration) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("Function\tRun\tDuration(ns)\n")
	if err != nil {
		return err
	}

	for name, durations := range data {
		for i, duration := range durations {
			_, err := file.WriteString(fmt.Sprintf("%s\t%d\t%d\n", name, i+1, duration.Nanoseconds()))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// O(1)

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

// O(log(n))

func binStr(n int) bool {
	binaryStr := strconv.FormatInt(int64(n), 2)
	return binaryStr[len(binaryStr)-1] == '0'
}

func regex(n int) bool {
	match, _ := regexp.MatchString(`.*[02468]$`, strconv.Itoa(n))

	return match
}

func str(n int) bool {
	numStr := strconv.Itoa(n)
	lastChar := numStr[len(numStr)-1]

	return lastChar == '0' || lastChar == '2' || lastChar == '4' || lastChar == '6' || lastChar == '8'
}

// O(n)

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
