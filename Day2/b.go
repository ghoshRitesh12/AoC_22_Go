package Day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Q https://adventofcode.com/2022/day/2#part2

func getAllMaps() (map[string] int, map[string] string, map[string] int, map[string] string) {
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
		"X": 1, // rock
		"Y": 2, // paper
		"Z": 3, // scissor
	}
	conditionalWin := map[string] string {
		"AX": "Z",
		"AY": "X",
		"AZ": "Y",

		"BX": "X",
		"BY": "Y",
		"BZ": "Z",

		"CX": "Y",
		"CY": "Z",
		"CZ": "X",
	}

	return outcome, lettersMap, scoresMap, conditionalWin
}

func B(filepath string) {
	file, err := os.Open(filepath);
	if err != nil {
		fmt.Println("Error: ", err);
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file);
	if scannerErr := scanner.Err(); scannerErr != nil {
		fmt.Println("Scanner error: ", scannerErr)
		return
	}

	totalScore := 0
	outcome, lettersMap, scoresMap, conditionalWin := getAllMaps()

	for scanner.Scan() {
		splittedInput := strings.Split(scanner.Text(), " ")
		elfPick := splittedInput[0]
		myPick := splittedInput[1]

		conditionWinArg := conditionalWin[elfPick+myPick]

		shapeScore := scoresMap[conditionWinArg]
		roundScore := outcome[lettersMap[elfPick] + " " + lettersMap[conditionWinArg]]
		
		totalScore += (shapeScore + roundScore)
	}

	fmt.Println("")
	fmt.Println("Day 2 > Q.b")
	fmt.Println("Total Score: ", totalScore);
}
