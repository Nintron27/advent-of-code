package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// totals := getInput("input_test")

	// fmt.Println("Part 1:", part1(totals))
	fmt.Println("Part 2:", part2())
}

func getInput(fileName string) (totals [][]int) {
	input, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")

	for _, line := range lines {
		temp := []int{}

		for _, char := range line {
			if char >= 48 && char <= 57 {
				temp = append(temp, int(char)-48)
			}
		}

		totals = append(totals, []int{temp[0], temp[len(temp)-1]})
	}

	return
}

func part1(input [][]int) int {
	lines := []int{}

	for _, val := range input {
		lines = append(lines, val[0]*10+val[1])
	}

	sum := 0
	for _, line := range lines {
		sum += line
	}

	return sum
}

func part2() int {
	file, err := os.Open("input")
	// input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scnr := bufio.NewScanner(file)

	for scnr.Scan() {
		fmt.Println(scnr.Text())
	}

	type valuePair struct {
		
	}

	var numMatches = []

	// lines := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")

	// numMap := map[string]int{
	// 	"one":   1,
	// 	"two":   2,
	// 	"three": 3,
	// 	"four":  4,
	// 	"five":  5,
	// 	"six":   6,
	// 	"seven": 7,
	// 	"eight": 8,
	// 	"nine":  9,
	// 	"1":     1,
	// 	"2":     2,
	// 	"3":     3,
	// 	"4":     4,
	// 	"5":     5,
	// 	"6":     6,
	// 	"7":     7,
	// 	"8":     8,
	// 	"9":     9,
	// }

	// type pair struct {
	// 	index int
	// 	val   int
	// }

	// totals := [][]int{}
	// for _, line := range lines {
	// 	pairs := []pair{}

	// 	for key, item := range numMap {
	// 		index := strings.Index(line, key)

	// 		if index != -1 {
	// 			pairs = append(pairs, pair{index: index, val: item})
	// 		}

	// 		index = strings.LastIndex(line, key)

	// 		if index != -1 {
	// 			pairs = append(pairs, pair{index: index, val: item})
	// 		}
	// 	}

	// 	sort.SliceStable(pairs, func(i, j int) bool {
	// 		return pairs[i].index < pairs[j].index
	// 	})

	// 	totals = append(totals, []int{pairs[0].val, pairs[len(pairs)-1].val})
	// }

	// stuff := []int{}
	// for _, val := range totals {
	// 	stuff = append(stuff, val[0]*10+val[1])
	// }

	// sum := 0
	// for _, line := range stuff {
	// 	sum += line
	// }

	return 42
}
