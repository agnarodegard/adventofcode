package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	content, err := os.ReadFile("2023/01-input.txt")
	if err != nil {
		log.Fatal("os.ReadFile error:", err)
	}

	lines := strings.Split(string(content), "\n")

	lines = []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	lines = []string{"two1nine"}

	//part1(lines)
	part2(lines)

}

func part2(lines []string) {

	sum := 0

	for _, line := range lines {

		// small number before big number is ok
		// one + eight
		// three + eight
		// five + eight

		// big number before small number is bad
		// nine + eight
		// eight + two
		// eight + three
		// two + one

		// Removing t in eight because three, two could be next
		line = strings.ReplaceAll(line, "nineigh", "9")

		// Removing o in two because one could be next
		line = strings.ReplaceAll(line, "eightwone", "81")

		// Removing o in two because one could be next
		line = strings.ReplaceAll(line, "eightw", "8")

		// Removing e in three, because eight could be next
		line = strings.ReplaceAll(line, "eighthr", "8")

		// Removing e in one because one could be next
		line = strings.ReplaceAll(line, "twoni", "2")

		// Removing e in one because one could be next
		line = strings.ReplaceAll(line, "twon", "2")

		line = strings.ReplaceAll(line, "one", "1")
		line = strings.ReplaceAll(line, "two", "2")
		line = strings.ReplaceAll(line, "three", "3")
		line = strings.ReplaceAll(line, "four", "4")
		line = strings.ReplaceAll(line, "five", "5")
		line = strings.ReplaceAll(line, "six", "6")
		line = strings.ReplaceAll(line, "seven", "7")
		line = strings.ReplaceAll(line, "eight", "8")
		line = strings.ReplaceAll(line, "nine", "9")
		sum += parseLine(line)
	}
	fmt.Println(sum)

}

func part1(lines []string) {
	sum := 0

	for _, line := range lines {
		sum += parseLine(line)
	}

	fmt.Println(sum)
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

	number, err := strconv.Atoi(b.String())
	if err != nil {
		log.Fatal("strconv.Atoi error:", err)
	}
	fmt.Println(number)
	return number
}
