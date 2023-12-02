package Day6

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func hasDuplicateChar(src []byte) bool {
	uniqueSrc := []byte{}
	hasCond := false
	
	for _, v := range src {
		if !slices.Contains[[]byte, byte](uniqueSrc, v) {
			uniqueSrc = append(uniqueSrc, v)
		} else {
			hasCond = true
			break
		}
	}

	return hasCond
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

	fmt.Println()
	fmt.Println("Day 6 > Q.a")

	if scanner.Scan() {
		text := []byte(scanner.Text())
		length := len(text)

		for i := 0; i < length; i++ {
			if (i + 3) <= (length - 1) && !hasDuplicateChar([]byte{
				text[i],
				text[i + 1],
				text[i + 2],
				text[i + 3],
			}) {
				fmt.Println("Characters processed before start-of-packet marker: ", (i+3) + 1)
				break
			}
		}
	}
}