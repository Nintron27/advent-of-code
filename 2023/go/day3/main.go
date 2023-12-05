package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	// lines := getInputLines("input")
	// fmt.Println("Part 1:", part1(lines))
	// fmt.Println("Part 2:", part2(lines))

	speed()
}

func speed() {
	yayaya := make([]int, 1)
	small := make([]int, 500000)
	medium := make([]int, 5000000)
	large := make([]int, 50000000)
	xlarge := make([]int, 500000000)

	start := time.Now()
	yayaya = append(yayaya, small...)
	fmt.Println("Small:", time.Since(start))

	start = time.Now()
	yayaya = append(yayaya, medium...)
	fmt.Println("Medium:", time.Since(start))

	start = time.Now()
	yayaya = append(yayaya, large...)
	fmt.Println("large:", time.Since(start))

	start = time.Now()
	yayaya = append(yayaya, xlarge...)
	fmt.Println("xlarge:", time.Since(start))
}

func getInputLines(fileName string) []string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")
}

func part1(lines []string) (sum int) {
	// parse out numbers

	for lsi, line := range lines {
		buildingInt := false
		num := 0
		digits := 0

		for li, c := range line {
			if c >= '0' && c <= '9' {
				buildingInt = true
			}

			if buildingInt {
				if !(c >= '0' && c <= '9') {
					buildingInt = false

					// NOW FIND IF YOU ADD!
					valid := false
					if lsi != 0 {
						if checkForSymbol(lines[lsi-1][slices.Max([]int{li - digits - 1, 0}):slices.Min([]int{li + 1, len(line)})]) {
							valid = true
						}
					}

					if checkForSymbol(line[slices.Max([]int{li - digits - 1, 0}):slices.Min([]int{li + 1, len(line)})]) {
						valid = true
					}

					if lsi < len(lines)-1 {
						if checkForSymbol(lines[lsi+1][slices.Max([]int{li - digits - 1, 0}):slices.Min([]int{li + 1, len(line)})]) {
							valid = true
						}
					}

					if valid {
						sum += num
					}

					num = 0
					digits = 0

				} else {
					digits++
					num *= 10
					thingy, _ := strconv.Atoi(string(c))
					num += thingy
				}
			}

			if li == len(line)-1 {
				valid := false
				if lsi != 0 {
					if checkForSymbol(lines[lsi-1][slices.Max([]int{li - digits - 1, 0}):slices.Min([]int{li + 1, len(line)})]) {
						valid = true
					}
				}

				if checkForSymbol(line[slices.Max([]int{li - digits - 1, 0}):slices.Min([]int{li + 1, len(line)})]) {
					valid = true
				}

				if lsi < len(lines)-1 {
					if checkForSymbol(lines[lsi+1][slices.Max([]int{li - digits - 1, 0}):slices.Min([]int{li + 1, len(line)})]) {
						valid = true
					}
				}

				if valid {
					sum += num
				}

				num = 0
				digits = 0
			}
		}
	}

	return
}

func part2(lines []string) (sum int) {
	// loop through everything, on asterisk, find numbers around
	// If there is exactly 2 numbers, then go "sideways" on those
	// to find them, and multiply them together, add to sum

	for li, line := range lines {
		for ri, r := range line {
			if r == '*' {
				numCount := 0
				numPos := make([][2]int, 0, 6)

				if li != 0 {
					num, pos := findSeparateInts(lines[li-1][max(ri-1, 0):min(ri+2, len(line))], li-1, max(ri-1, 0))

					numCount += num
					numPos = append(numPos, pos...)
				}

				num, pos := findSeparateInts(line[max(ri-1, 0):min(ri+2, len(line))], li, max(ri-1, 0))

				numCount += num
				numPos = append(numPos, pos...)

				if li < len(lines)-1 {
					num, pos := findSeparateInts(lines[li+1][max(ri-1, 0):min(ri+2, len(line))], li+1, max(ri-1, 0))

					numCount += num
					numPos = append(numPos, pos...)
				}

				if numCount == 2 {
					// Now, build the numbers
					tempSum := 1

				pairLoop:
					for _, pair := range numPos {
						buildingNum := false
						num := 0
						intersecting := false
						for ri, r := range lines[pair[0]] {
							if r >= '0' && r <= '9' {
								buildingNum = true
								if ri == pair[1] {
									intersecting = true
								}
							}

							if buildingNum {
								if !(r >= '0' && r <= '9') {
									if intersecting {
										tempSum *= num
										continue pairLoop
									}
									buildingNum = false
									intersecting = false
									num = 0
								} else {
									num *= 10
									int, _ := strconv.Atoi(string(r))
									num += int
								}
							}

							if ri == len(lines[pair[0]])-1 {
								tempSum *= num
							}
						}
					}

					sum += tempSum
				}
			}
		}
	}

	return
}

func findSeparateInts(input string, lineIndex, inputIndex int) (count int, positions [][2]int) {
	sameNum := false

	for i, c := range input {
		if c >= '0' && c <= '9' {
			if !sameNum {
				count++
				sameNum = true

				positions = append(positions, [2]int{lineIndex, inputIndex + i})
			}
		} else {
			sameNum = false
		}
	}

	return
}

func checkForSymbol(input string) bool {
	for _, c := range input {
		if isSymbol(c) {
			return true
		}
	}

	return false
}

func isSymbol(r rune) bool {
	notNum := !(r >= '0' && r <= '9')
	return notNum && r != '.'
}

type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func min[T Numeric](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T Numeric](a, b T) T {
	if a > b {
		return a
	}
	return b
}
