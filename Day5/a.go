package Day5

import (
	"bufio"
	"fmt"
	"os"
)

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

	stacks := make([]stack, 9)
	indexRoundBackNo := 4

	scanner.Scan()
	for (scanner.Text() != " 1   2   3   4   5   6   7   8   9 ") {
		for i, v := range scanner.Text() {
			if v != ' ' && v != '[' && v != ']' {
				stacks[i/indexRoundBackNo].Push(string(v))
			}
		}
		scanner.Scan()
	}

	// read the empty line
	scanner.Scan()

	for scanner.Scan() {
		var moveAmount, from, to int;
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &moveAmount, &from, &to)

		for i := 0; i < moveAmount; i++ {
			stacks[to-1].Push(stacks[from-1].Pop())
		}
	}

	var topCrates string
	for _, s := range stacks {
		topCrates += s.Peek()
	}

	fmt.Println()
	fmt.Println("Day 5 > Q.a")
	fmt.Println("Top crates: ", topCrates)
}