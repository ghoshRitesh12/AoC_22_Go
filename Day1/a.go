package D1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Q: https://adventofcode.com/2022/day/1

func A(filepath string) {	
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error: ", err);
		return;
	}

	defer file.Close()

	sum := 0
	totalElfCalorie := []int{}

	scanner := bufio.NewScanner(file);

	for scanner.Scan() {
		line := scanner.Text()

		val, _ := strconv.Atoi(line)
		sum += val

		comparisonResult := strings.Compare(strings.TrimSpace(line), "")
		if comparisonResult == 0 {
			totalElfCalorie = append(totalElfCalorie, sum);
			sum = 0
		}
	}
	
	max := totalElfCalorie[0]
	for i := 1; i < len(totalElfCalorie); i++ {
		if(totalElfCalorie[i] > max) {
			max = totalElfCalorie[i]
		}
	}

	fmt.Println("")
	fmt.Println("Day 1 > Qa")
	fmt.Println("Max amount is: ", max);

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
	}
}
