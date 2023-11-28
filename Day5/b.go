package Day5

import (
	"bufio"
	"fmt"
	"os"
)

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

	stacks := make([]stack, 9)
	indexRoundBackNo := 4

	scanner.Scan()
	for (scanner.Text() != " 1   2   3   4   5   6   7   8   9 ") {
		for i, v := range scanner.Text() {
			if v != ' ' && v != '[' && v != ']' {
				stacks[i/indexRoundBackNo].AddToBottom(string(v))
			}
		}
		scanner.Scan()
	}

	// read the empty line
	scanner.Scan()

	for scanner.Scan() {
		var moveAmount, from, to int;
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &moveAmount, &from, &to)

		popItems := []string{}

		for i := 0; i < moveAmount; i++ {
			poppedItem := stacks[from-1].Pop()
			popItems = append([]string{poppedItem}, popItems...)
		}
		
		for i := 0; i < len(popItems); i++ {
			stacks[to-1].Push(popItems[i])
		}
	}

	var topCrates string
	for _, s := range stacks {
		topCrates += s.Peek()
	}

	fmt.Println()
	fmt.Println("Day 5 > Q.b")
	fmt.Println("Top crates: ", topCrates)
}
