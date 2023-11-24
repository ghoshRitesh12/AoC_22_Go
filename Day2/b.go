package Day2

import (
	"bufio"
	"fmt"
	"os"
)

// Q https://adventofcode.com/2022/day/2#part2

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
	scores := map[string] int {
		"B X": 1, 
		"C X": 2, 
		"A X": 3, 
		"A Y": 4,
		"B Y": 5,
		"C Y": 6,
		"C Z": 7,
		"A Z": 8,
		"B Z": 9,
	}
	totalScore := 0

	for scanner.Scan() {		
		totalScore += scores[scanner.Text()]	
	}

	fmt.Println("")
	fmt.Println("Day 2 > Q.b")
	fmt.Println("Total Score: ", totalScore);
}
