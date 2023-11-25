package Day3

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"unicode"
)

// Q https://adventofcode.com/2022/day/3

func getLetterPriority(letterAscii int) int {
	priority := 0

	if unicode.IsLower(rune(letterAscii)) {
		priority = letterAscii - 96
	} else {
		priority = letterAscii - 38
	}

	return priority
}

func findDuplicateItem(src1 []byte, src2 []byte) int {
	var duplicateItem int

	for _, value := range src1 {
		val := byte(value)
		if slices.Contains[[]byte, byte](src2, val) {
			duplicateItem = int(value)
			break
		}
	}

	return duplicateItem
}

func A(filepath string) {
	file, err := os.Open(filepath);
	if err != nil {
		fmt.Println("Error: ", err);
		return;
	}
	defer file.Close()

	scanner := bufio.NewScanner(file);
	if scannerErr := scanner.Err(); scannerErr != nil {
		fmt.Println("Scanner error: ", scannerErr)
		return
	}

	totalPriorities := 0

	for scanner.Scan() {
		scannedString := scanner.Text()
		stringLength := len(scannedString)
		halfLength := stringLength / 2

		firstHalf := scannedString[0 : halfLength]
		secondHalf := scannedString[halfLength : stringLength]

		duplicateItem := findDuplicateItem([]byte(firstHalf), []byte(secondHalf))

		totalPriorities += getLetterPriority(duplicateItem)
	}

	fmt.Println("")
	fmt.Println("Day 3 > Q.a")
	fmt.Println("Total Priorities: ", totalPriorities);
}
