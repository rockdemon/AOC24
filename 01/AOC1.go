package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Load the input file
	data, err := os.ReadFile("input")
	if err != nil {
		return
	}

	// Split the input into lines
	inp := strings.Split(string(data), "\n")

	var leftdigits []int
	var rightdigits []int
	var total int = 0
	// Process each line
	for _, s := range inp {
		strarrRow := strings.Fields(s)
		if len(strarrRow) < 2 {
			continue
		}

		left, err := strconv.Atoi(strarrRow[0])
		right, err := strconv.Atoi(strarrRow[1])

		leftdigits = append(leftdigits, left)
		rightdigits = append(rightdigits, right)
	}

	// Use leftdigits and rightdigits as needed
	fmt.Println("Left digits:", leftdigits)
	fmt.Println("Right digits:", rightdigits)
	sort.Ints(leftdigits)
	sort.Ints(rightdigits)
	fmt.Println("Left digits:", leftdigits)
	fmt.Println("Right digits:", rightdigits)
	for i := 0; i < len(leftdigits); i++ {
		diff := leftdigits[i] - rightdigits[i]
		if diff > 0 {
			total += diff
		} else {
			total -= diff
		}

	}
	fmt.Println("Total:", total)
}
