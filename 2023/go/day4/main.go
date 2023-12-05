package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	lines := getInputLines("input")
	// fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
	fmt.Println("Time:", time.Since(start))
}

func part1(lines []string) (sum int) {
	for _, line := range lines {
		input := strings.SplitN(line, ": ", 2)[1]
		// input = strings.Trim(input, " ")

		// fmt.Println(input)

		split := strings.SplitN(input, " | ", 2)
		winnings := strings.Split(split[0], " ")
		mine := strings.Split(split[1], " ")

		yaya := 1

		for _, winNum := range winnings {
			for _, myNum := range mine {
				if myNum == winNum && len(myNum) > 0 {
					fmt.Println(myNum, winNum)
					yaya = yaya << 1
				}
			}
		}

		// fmt.Println("line:", i+1)
		// fmt.Println("winnings:", yaya>>1)

		sum += yaya >> 1
	}

	return
}

func part2(lines []string) (sum int) {
	copies := make(map[int]int)

	for li, line := range lines {
		input := strings.SplitN(line, ": ", 2)[1]

		split := strings.SplitN(input, " | ", 2)
		winnings := strings.Split(split[0], " ")
		mine := strings.Split(split[1], " ")

		winningNums := 0

		for _, winNum := range winnings {
			for _, myNum := range mine {
				if myNum == winNum && len(myNum) > 0 {
					winningNums++
				}
			}
		}

		copies[li]++

		for i := 1; i <= winningNums; i++ {
			copies[li+i] += 1 * copies[li]
		}

		if winningNums == 0 {
			winningNums = 1
		}

		sum += copies[li]
	}

	return
}

func getInputLines(fileName string) []string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")
}
