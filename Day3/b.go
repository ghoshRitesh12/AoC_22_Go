package Day3

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// Q https://adventofcode.com/2022/day/3#part2

func getUniqueSlice(inputSlice []byte) []byte {
	uniqueSlice := []byte{}
	for _, value := range inputSlice {
		if !slices.Contains[[]byte, byte](uniqueSlice, value) {
			uniqueSlice = append(uniqueSlice, value)
		}
	}
	return uniqueSlice
}

func findTwoDuplicateItems(src1 []byte, src2 []byte, src3 []byte) int {
	var duplicateItem int
	duplicateCounter := 1

	for _, val := range src1 {
		if duplicateCounter == 3  {
			break
		}
		if slices.Contains[[]byte, byte](src2, val) && slices.Contains[[]byte, byte](src3, val) {
			duplicateItem = int(val)
			duplicateCounter += 1
		}
	}

	return duplicateItem
}

func B(filepath string) {
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
		firstHalf := getUniqueSlice([]byte(scanner.Text()))
		scanner.Scan()
		secondHalf := getUniqueSlice([]byte(scanner.Text()))
		scanner.Scan()
		thirdHalf := getUniqueSlice([]byte(scanner.Text()))
		
		duplicateItem := findTwoDuplicateItems(firstHalf, secondHalf, thirdHalf)

		totalPriorities += getLetterPriority(duplicateItem)
	}

	fmt.Println("")
	fmt.Println("Day 3 > Q.b")
	fmt.Println("Total Priorities: ", totalPriorities);
}