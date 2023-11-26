package Day4

import (
	"bufio"
	"fmt"
	"os"
)

// Q https://adventofcode.com/2022/day/4

func isRangeFullyContained(s1Num1, s1Num2, s2Num1, s2Num2 int) bool {
	if s1Num1 >= s2Num1 && s1Num2 <= s2Num2 || s1Num1 <= s2Num1 && s1Num2 >= s2Num2 {
		return true
	}

	return false
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

	sumOfAssignmentPairs := 0

	for scanner.Scan() {
		var firstSrcNum1, firstSrcNum2, secondSrcNum1, secondSrcNum2 int
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &firstSrcNum1, &firstSrcNum2, &secondSrcNum1, &secondSrcNum2)

		if isRangeFullyContained(firstSrcNum1, firstSrcNum2, secondSrcNum1, secondSrcNum2) {
			sumOfAssignmentPairs += 1
		}
	}

	fmt.Println("")
	fmt.Println("Day 4 > Q.a")
	fmt.Println("Total Assignment Pairs: ", sumOfAssignmentPairs);
}
