package Day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Q: https://adventofcode.com/2022/day/1#part2

func findLargestNumbers(numberSlice []int) (int, int, int) {
	sort.Sort(sort.Reverse(sort.IntSlice(numberSlice)))
	return numberSlice[0], numberSlice[1], numberSlice[2]
}

func B(filepath string) {	
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error: ", err);
		return;
	}

	defer file.Close()

	sum := 0
	totalElfCalorie := []int{}

	scanner := bufio.NewScanner(file);
	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
	}

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
	
	one, two, three := findLargestNumbers(totalElfCalorie)	
	largestNumbers := [...]int{one, two, three}
	
	sum = 0
	for i := 0; i<len(largestNumbers); i++ {
		sum += largestNumbers[i]
	}

	fmt.Println("")
	fmt.Println("Day 1 > Q.b")
	fmt.Println("Sum of calories carried by top 3 elves are: ", sum);
}
