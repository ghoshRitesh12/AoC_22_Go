package Day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Q https://adventofcode.com/2022/day/2

func getMaps() (map[string] int, map[string] string, map[string] int) {
	outcome := map[string] int {
		"rock rock": 3,
		"rock paper": 6,
		"rock scissor": 0,

		"paper paper": 3,
		"paper rock": 0,
		"paper scissor": 6,

		"scissor scissor": 3,
		"scissor paper": 0,
		"scissor rock": 6,
	}
	lettersMap := map[string] string {
		"A": "rock",
		"B": "paper",
		"C": "scissor",

		"X": "rock",
		"Y": "paper",
		"Z": "scissor",
	}
	scoresMap := map[string] int {
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	return outcome, lettersMap, scoresMap
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

	totalScore := 0
	outcome, lettersMap, scoresMap := getMaps()

	for scanner.Scan() {
		splittedInput := strings.Split(scanner.Text(), " ")
		elfPick := splittedInput[0]
		myPick := splittedInput[1]

		shapeScore := scoresMap[myPick]
		roundScore := outcome[lettersMap[elfPick] + " " + lettersMap[myPick]]
		
		totalScore = totalScore + (shapeScore + roundScore)
	}

	fmt.Println("")
	fmt.Println("Day 2 > Q.a")
	fmt.Println("Total Score: ", totalScore);
}
