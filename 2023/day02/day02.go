package day02

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	gameID int
	css    []cubeSet
	power  int
}

type cubeSet struct {
	red   int
	green int
	blue  int
}

func Day02() {

	content, err := os.ReadFile("2023/day02/day02-input.txt")
	if err != nil {
		log.Fatal("os.ReadFile error:", err)
	}

	lines := strings.Split(string(content), "\n")

	restriction := cubeSet{
		red:   12,
		green: 13,
		blue:  14,
	}
	validGamesSum, power := day02part1(lines, restriction)
	fmt.Printf("Day 01: part 1: %d part 2:%d\n", validGamesSum, power)

}

func day02part1(lines []string, r cubeSet) (int, int) {

	gameIDsum := 0
	power := 0

	for _, line := range lines {
		parsedGame := parseGame(line)

		if validateGame(parsedGame, r) {
			gameIDsum += parsedGame.gameID
		}
		power += parsedGame.power

	}
	return gameIDsum, power
}

func parseGame(line string) (g game) {
	game := strings.Split(line, ":")
	reID := regexp.MustCompile("[^0-9]+")

	foundIDString := reID.ReplaceAllString(game[0], "")
	foundIDInt, err := strconv.Atoi(foundIDString)
	if err != nil {
		log.Fatal("strconv.Atoi error: ", err)
	}

	g.gameID = foundIDInt
	g.css = parseGameSets(game[1])
	smallestRestriction := findSmallestRestriction(g.css)
	g.power = sumColors(smallestRestriction)

	return g
}

func sumColors(cs cubeSet) int {
	return cs.red * cs.green * cs.blue
}

func validateGame(g game, restriction cubeSet) bool {

	valid := true

	for _, gameSet := range g.css {

		if !valid {
			continue
		}

		valid = validateGameSet(gameSet, restriction)

	}

	return valid
}

func validateGameSet(cs cubeSet, r cubeSet) bool {
	return cs.red <= r.red && cs.green <= r.green && cs.blue <= r.blue
}

func parseGameSets(line string) (css []cubeSet) {

	sets := strings.Split(line, ";")
	for _, set := range sets {
		css = append(css, parseSet(set))
	}
	return css

}

func findSmallestRestriction(css []cubeSet) (smcs cubeSet) {

	for _, cs := range css {
		if smcs.red < cs.red {
			smcs.red = cs.red
		}
		if smcs.green < cs.green {
			smcs.green = cs.green
		}
		if smcs.blue < cs.blue {
			smcs.blue = cs.blue
		}
	}
	return smcs
}

// parseSet takes a game set and returns a cubeSet struct with the correct number of colored cubes.
func parseSet(line string) (cs cubeSet) {

	colors := strings.Split(line, ",")
	reColor := regexp.MustCompile("[^a-zA-Z]+")
	reAmount := regexp.MustCompile("[^0-9]+")
	for _, color := range colors {
		foundColor := reColor.ReplaceAllString(color, "")
		foundAmountString := reAmount.ReplaceAllString(color, "")
		foundAmountInt, err := strconv.Atoi(foundAmountString)
		if err != nil {
			log.Fatal("strconv.Atoi error: ", err)
		}

		switch foundColor {
		case "red":
			cs.red = foundAmountInt
		case "green":
			cs.green = foundAmountInt
		case "blue":
			cs.blue = foundAmountInt
		}

	}

	return cs
}
