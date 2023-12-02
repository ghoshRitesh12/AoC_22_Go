package Day7

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	dirnamesHistory := []string{}
	directories := make(map[string]int)

	for scanner.Scan() {
		text := scanner.Text()
		var __dirname string
		var fileSpace int

		fmt.Sscanf(text, "$ cd %s", &__dirname)
		fmt.Sscanf(text, "%d ", &fileSpace)

		
		if text[:4] == "$ cd" && __dirname != "" {
			if strings.TrimSpace(__dirname) == ".." {
				dirnamesHistory = dirnamesHistory[:len(dirnamesHistory) - 1]
			} else {
				dirnamesHistory = append(dirnamesHistory, __dirname)
			}
		} 
		
		if fileSpace != 0 {
			for i := 0; i < len(dirnamesHistory); i++ {
				key := strings.Join(dirnamesHistory[:i+1], "/")
				directories[key] += fileSpace 
			}
		}
	}

	maxDirSpace := directories["/"]
	minDirSpaces := []int{}

	unusedSpace := 70000000 - maxDirSpace
	requiredSpace := 30000000 - unusedSpace

	for _, directorySpace := range directories {
		if directorySpace >= requiredSpace {
			minDirSpaces = append(minDirSpaces, directorySpace)
		}
	}

	minDirSpace := minDirSpaces[0]
	for i := 0; i < len(minDirSpaces); i++ {
		if(minDirSpaces[i] < minDirSpace) {
			minDirSpace = minDirSpaces[i]
		}
	}
	
	fmt.Println()
	fmt.Println("Day 7 > Q.b")
	fmt.Println("Min Dir Space to free up: ", minDirSpace)
}
