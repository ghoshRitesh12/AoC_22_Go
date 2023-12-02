package Day6

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

	fmt.Println()
	fmt.Println("Day 6 > Q.b")

	if scanner.Scan() {
		text := []byte(scanner.Text())
		length := len(text)

		for i := 0; i < length; i++ {
			if (i + 13) <= (length - 1) && !hasDuplicateChar([]byte{
				text[i],
				text[i + 1],
				text[i + 2],
				text[i + 3],
				text[i + 4],
				text[i + 5],
				text[i + 6],
				text[i + 7],
				text[i + 8],
				text[i + 9],
				text[i + 10],
				text[i + 11],
				text[i + 12],
				text[i + 13],
			}) {
				fmt.Println("Characters processed before start-of-message marker: ", (i+13) + 1)
				break
			}
		}
	}
}