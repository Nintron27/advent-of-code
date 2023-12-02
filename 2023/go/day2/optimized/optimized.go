package optimized

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1(fileName string) (sum int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		// Split on first ": " to separate the game id and the cube sets
		gameSplitStr := strings.SplitN(scanner.Text(), ": ", 2)
		id, _ := strconv.Atoi(gameSplitStr[0][5:])
		cubeSets := strings.Split(gameSplitStr[1], "; ")

		for _, cubeSet := range cubeSets {
			// Split on every ", " to get the cube and it's count
			cubesStr := strings.Split(cubeSet, ", ")

			// For every cube, go rune by rune, adding up the number
			// Once a space is found, grab the next rune as the color
			// since red, green, and blue start with different letters
			// we only need one rune to identify them
			for _, cube := range cubesStr {
				// red, green, and blue start with different letters
				// so we only need one rune to identify them
				var color rune
				var num uint8
				for i, c := range cube {
					if c != ' ' {
						// num can only be 2 digits
						// so on the first loop its 0 * 10
						// but on the second loop it's moving the first
						// loops digit to the 10s place
						num *= 10
						num += uint8(c - 48) // convert ASCII into int
					} else {
						color = rune(cube[i+1])
						break
					}
				}

				if num > 14 ||
					// find the colors by their 0th index in the alphabet
					color == 'r' && num > 12 ||
					color == 'g' && num > 13 ||
					color == 'b' && num > 14 {
					// Skip checking the rest of the entire game
					// if one color is too much
					goto endCubeSet
				}
			}
		}

		sum += id
	endCubeSet:
	}

	return
}

// func part2(lines []string) (sum int) {
// 	for _, line := range lines {
// 		// id, _ := strconv.Atoi(strings.Split(line[5:], ": ")[0])

// 		cubeSets := strings.Split(strings.Split(line, ":")[1], ";")

// 		minSet := make(map[string]int)
// 		for _, cubeSet := range cubeSets {
// 			cubeMap := make(map[string]int)

// 			cubes := strings.Split(cubeSet, ",")

// 			for _, cube := range cubes {
// 				cube = strings.Trim(cube, " ")

// 				cubeSplit := strings.Split(cube, " ")

// 				num, _ := strconv.Atoi(cubeSplit[0])

// 				cubeMap[cubeSplit[1]] = num
// 			}

// 			for k, _ := range cubeMap {
// 				if cubeMap[k] > minSet[k] {
// 					minSet[k] = cubeMap[k]
// 				}
// 			}
// 		}

// 		powerSum := 1

// 		for k, _ := range minSet {
// 			powerSum *= minSet[k]
// 		}

// 		sum += powerSum
// 	}

// 	return
// }
