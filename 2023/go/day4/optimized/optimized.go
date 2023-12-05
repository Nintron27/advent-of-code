package optimized

import (
	"log"
	"os"
	"strings"
)

func getInputLines(fileName string) []string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")
}

func part2(lines []string) (sum int) {
	cardCopies := make([]int, len(lines))

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

		cardCopies[li]++

		for i := 1; i <= winningNums; i++ {
			cardCopies[li+i] += 1 * cardCopies[li]
		}

		if winningNums == 0 {
			winningNums = 1
		}

		sum += cardCopies[li]
	}

	return
}
