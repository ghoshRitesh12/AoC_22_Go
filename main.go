package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ghoshRitesh12/AoC_22_Go/Day1"
	"github.com/ghoshRitesh12/AoC_22_Go/Day2"
	"github.com/ghoshRitesh12/AoC_22_Go/Day3"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Choose a day to execute the problem solution: ")
	inputDay, err := reader.ReadString('\n');
	if err != nil {
		fmt.Println("Opps an error occured")
		os.Exit(1)
	}

	day, err := strconv.Atoi(strings.TrimSpace(inputDay))

	if err != nil {
		fmt.Println("Opps an error occured")
		os.Exit(1)
	}
	if day <= 0 || day > 25 {
		fmt.Println("The day entered is invalid")
		os.Exit(1)
	}

	fmt.Println("Choose a problem no to execute in lowercase (a or b): ")
	inputProblemNo, err := reader.ReadString('\n');
	if err != nil {
		fmt.Println("Opps an error occured")
		os.Exit(1)
	}

	problemNo := strings.TrimSpace(inputProblemNo)

	if problemNo != "a" && problemNo != "b" {
		fmt.Println("The problem no entered is invalid")
		os.Exit(1)
	}

	switch(day) {
		case 1: {
			if problemNo == "a" {
				Day1.A("./Day2/input.txt")

			} else if problemNo == "b" {
				Day1.B("./Day2/input.txt")

			}
		}
		case 2: {
			if problemNo == "a" {
				Day2.A("./Day2/input.txt")

			} else if problemNo == "b" {
				Day2.B("./Day2/input.txt")

			}
		}
		case 3: {
			if problemNo == "a" {
				Day3.A("./Day3/input.txt")

			} else if problemNo == "b" {
				Day2.B("./Day2/input.txt")

			}
		}

	}
}
