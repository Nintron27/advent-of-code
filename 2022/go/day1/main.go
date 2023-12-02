package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	totals := getInput("input")

	fmt.Println("Part 1:", part1(totals))
	fmt.Println("Part 2:", part2(totals))
}

func getInput(fileName string) (totals []int) {
	inputBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	inputStr := strings.TrimSuffix(string(inputBytes), "\n")
	inventories := strings.Split(inputStr, "\n\n")

	for _, inventory := range inventories {
		calItems := strings.Split(inventory, "\n")

		sum := 0

		for _, item := range calItems {
			num, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal(err)
			}

			sum += num
		}

		totals = append(totals, sum)
	}

	return
}

func part1(totals []int) int {
	return slices.Max(totals)
}

func part2(totals []int) int {
	slices.Sort(totals)

	sum := 0
	for _, num := range totals[len(totals)-3:] {
		sum += num
	}

	return sum
}
