package main

import (
	"day2/optimized"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	lines := getInputLines("input")
	fmt.Println("Part 1:", part1(lines))
	fmt.Printf("Time: %v\n", time.Since(start))
	// fmt.Println("Part 2:", part2(lines))

	startOptimized := time.Now()
	fmt.Println("Part 1 opt:", optimized.Part1("input"))
	fmt.Printf("Time opt: %v\n", time.Since(startOptimized))
}

func getInputLines(fileName string) []string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")
}

func part1(lines []string) (sum int) {
	for _, line := range lines {
		id, _ := strconv.Atoi(strings.Split(line[5:], ": ")[0])

		cubeSets := strings.Split(strings.Split(line, ":")[1], ";")

		valid := true
		for _, cubeSet := range cubeSets {
			cubeMap := make(map[string]int)

			cubes := strings.Split(cubeSet, ",")

			for _, cube := range cubes {
				cube = strings.Trim(cube, " ")

				cubeSplit := strings.Split(cube, " ")

				num, _ := strconv.Atoi(cubeSplit[0])

				cubeMap[cubeSplit[1]] = num
			}

			if cubeMap["red"] > 12 || cubeMap["green"] > 13 || cubeMap["blue"] > 14 {
				valid = false
			}
		}

		if valid {
			sum += id
		}
	}

	return
}

func part2(lines []string) (sum int) {
	for _, line := range lines {
		// id, _ := strconv.Atoi(strings.Split(line[5:], ": ")[0])

		cubeSets := strings.Split(strings.Split(line, ":")[1], ";")

		minSet := make(map[string]int)
		for _, cubeSet := range cubeSets {
			cubeMap := make(map[string]int)

			cubes := strings.Split(cubeSet, ",")

			for _, cube := range cubes {
				cube = strings.Trim(cube, " ")

				cubeSplit := strings.Split(cube, " ")

				num, _ := strconv.Atoi(cubeSplit[0])

				cubeMap[cubeSplit[1]] = num
			}

			for k := range cubeMap {
				if cubeMap[k] > minSet[k] {
					minSet[k] = cubeMap[k]
				}
			}
		}

		powerSum := 1

		for k := range minSet {
			powerSum *= minSet[k]
		}

		sum += powerSum
	}

	return
}
