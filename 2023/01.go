package main

import (
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

	// pairs contains first and last number in each line, or a single number twice.
	pairs := []int{}

	for _, line := range lines {

		var runes = []int{}
		for _, r := range line {
			if r < 49 || r > 57 {
				continue
			}

			if len(runes) < 2 {
				runes = append(runes, int(r-'0'))
			} else {
				runes[1] = int(r - '0')
			}

		}

		pair := strings.Trim(strings.Join(strings.Fields(fmt.Sprintf("%d", runes)), ""), "[]")

		if len(pair) == 1 {
			pair = pair + pair
		}

		digits, err := strconv.Atoi(pair)
		if err != nil {
			log.Fatal("strconv.Atoi error:", err)
		}
		pairs = append(pairs, digits)

	}

	sum := 0

	for _, pair := range pairs {
		sum += pair
	}
	fmt.Println(sum)
}
