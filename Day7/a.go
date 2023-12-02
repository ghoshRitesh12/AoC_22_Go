package Day7

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	totalDirectorySpace := 0
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

	for _, directorySpace := range directories {
		if directorySpace <= 100000 {
			totalDirectorySpace += directorySpace
		}
	}
	
	fmt.Println()
	fmt.Println("Day 7 > Q.a")
	fmt.Println("Total Directory Space: ", totalDirectorySpace)
}
