package day01

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day01() {

	content, err := os.ReadFile("days/day01-input.txt")
	if err != nil {
		log.Fatal("os.ReadFile error:", err)
	}

	lines := strings.Split(string(content), "\n")

	//lines = []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	//lines = []string{"6r219sevenpcvfpmfxxl"}

	fmt.Printf("Solution part 1: %d\n", day01part1(lines))
	fmt.Printf("Solution part 1: %d\n", day01part2(lines))

}

func day01part2(lines []string) int {

	// sum holds the running sum over all lines.
	sum := 0

	// numberLookup is used to swap numberwords for numbernumbers
	numberLookup := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, line := range lines {

		/**
		Some numberwords may overlap, e.g. nineight, and finding a search order for numberwords that handles all cases is difficult and prone to errors
		Here we find the index of each numberword in the string, keeping the lowest index.
		Then we replace the first occurrence of numberword with numbernumber.
		*/

		// Start with a absurdly high lowestIndex to guarantee that any index we find is lower and thus allow us to record the match.
		lowestIndex := 10000
		var lowestMatch string

		for numberString := range numberLookup {
			foundIndex := strings.Index(line, numberString)
			if foundIndex == -1 {
				continue
			}

			if foundIndex < lowestIndex {
				lowestIndex = foundIndex
				lowestMatch = numberString
			}

			// It's tempting to also identify the highest index in the same loop, but overlapping numberwords may cause problems.
			// E.g. oneight will have lowestMatch one, and highestMatch eight, but when we swap one with 1, eight is no longer a valid highest match.
		}

		// Replace first matching numberWord with numberString. Safe if no match found.
		line = strings.Replace(line, lowestMatch, numberLookup[lowestMatch], 1)

		/**
		Having found the first numberword in the string, we now look for the last.
		*/

		// Start with the lowest possible index to guarantee that any index we find is higher and thus allow us to record the match.
		highestIndex := 0
		var highestMatch string

		for numberString := range numberLookup {
			foundIndex := strings.Index(line, numberString)
			if foundIndex == -1 {
				continue
			}

			if foundIndex > highestIndex {
				highestIndex = foundIndex
				highestMatch = numberString
			}
		}

		// Replace last matching numberWord with numberString. Safe if no match found.
		line = strings.Replace(line, highestMatch, numberLookup[highestMatch], 1)

		sum += parseLine(line)
	}

	return sum

}

func day01part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		sum += parseLine(line)
	}

	return sum
}

func parseLine(line string) int {
	var b bytes.Buffer
	for _, r := range line {
		if r < 49 || r > 57 {
			continue
		}

		if b.Len() > 1 {
			b.Truncate(1)
		}
		b.WriteRune(r)

	}

	if b.Len() == 1 {
		b.WriteString(b.String())
	}

	if b.Len() > 0 {
		number, err := strconv.Atoi(b.String())
		if err != nil {
			log.Fatal("strconv.Atoi error:", err)
		}
		return number
	}

	return 0

}
