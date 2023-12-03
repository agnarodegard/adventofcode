package day02

import (
	"testing"
)

func TestParseSet(t *testing.T) {

	// Set 1: 12 blue
	// Set 2: 2 green, 13 blue, 19 red

	type testCubeSetParse struct {
		input    string
		expected cubeSet
	}

	tests := []testCubeSetParse{
		{
			input: "12 blue",
			expected: cubeSet{
				red:   0,
				green: 0,
				blue:  12,
			},
		},
		{
			input: "2 green, 13 blue, 19 red",
			expected: cubeSet{
				red:   19,
				green: 2,
				blue:  13,
			},
		},
	}

	for _, tc := range tests {
		gotCubeSet := parseSet(tc.input)
		if gotCubeSet.red != tc.expected.red {
			t.Errorf("expected: %d red, got %d red", tc.expected.red, gotCubeSet.red)
		}
		if gotCubeSet.green != tc.expected.green {
			t.Errorf("expected: %d green, got %d green", tc.expected.green, gotCubeSet.green)
		}
		if gotCubeSet.blue != tc.expected.blue {
			t.Errorf("expected: %d blue, got %d blue", tc.expected.blue, gotCubeSet.blue)
		}
	}

}

func TestParseGameSets(t *testing.T) {

	// 12 blue; 2 green, 13 blue, 19 red; 13 red, 3 green, 14 blue

	type testCubeSetParse struct {
		input    string
		expected []cubeSet
	}

	tests := []testCubeSetParse{
		{
			input: "12 blue; 2 green, 13 blue, 19 red; 13 red, 3 green, 14 blue",
			expected: []cubeSet{
				{
					red:   0,
					green: 0,
					blue:  12,
				},
				{
					red:   19,
					green: 2,
					blue:  13,
				},
				{
					red:   13,
					green: 3,
					blue:  14,
				},
			},
		},
	}

	for _, tc := range tests {
		gotCubeSets := parseGameSets(tc.input)

		for k := range gotCubeSets {
			if gotCubeSets[k].red != tc.expected[k].red {
				t.Errorf("expected: %d red, got %d red", tc.expected[k].red, gotCubeSets[k].red)
			}
			if gotCubeSets[k].green != tc.expected[k].green {
				t.Errorf("expected: %d green, got %d green", tc.expected[k].green, gotCubeSets[k].green)
			}
			if gotCubeSets[k].blue != tc.expected[k].blue {
				t.Errorf("expected: %d blue, got %d blue", tc.expected[k].blue, gotCubeSets[k].blue)
			}
		}

	}

}

func TestParseGame(t *testing.T) {

	// Game 1: 12 blue; 2 green, 13 blue, 19 red; 13 red, 3 green, 14 blue

	type testCubeSetParse struct {
		input    string
		expected game
	}

	tests := []testCubeSetParse{
		{
			input: "Game 1: 12 blue; 2 green, 13 blue, 19 red; 13 red, 3 green, 14 blue",
			expected: game{
				gameID: 1,
				css: []cubeSet{
					{
						red:   0,
						green: 0,
						blue:  12,
					},
					{
						red:   19,
						green: 2,
						blue:  13,
					},
					{
						red:   13,
						green: 3,
						blue:  14,
					},
				},
			},
		},
		{
			input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expected: game{
				gameID: 3,
				css: []cubeSet{
					{
						red:   20,
						green: 8,
						blue:  6,
					},
					{
						red:   4,
						green: 13,
						blue:  5,
					},
					{
						red:   1,
						green: 5,
						blue:  0,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		gotGame := parseGame(tc.input)
		if gotGame.gameID != tc.expected.gameID {
			t.Errorf("expected: gameID %d, got %d", tc.expected.gameID, gotGame.gameID)
		}
		for k := range gotGame.css {
			if gotGame.css[k].red != tc.expected.css[k].red {
				t.Errorf("expected: %d red, got %d red", tc.expected.css[k].red, gotGame.css[k].red)
			}
			if gotGame.css[k].green != tc.expected.css[k].green {
				t.Errorf("expected: %d green, got %d green", tc.expected.css[k].green, gotGame.css[k].green)
			}
			if gotGame.css[k].blue != tc.expected.css[k].blue {
				t.Errorf("expected: %d blue, got %d blue", tc.expected.css[k].blue, gotGame.css[k].blue)
			}
		}

	}

}

func TestValidateGameSet(t *testing.T) {

	type testGameSetValidate struct {
		cs       cubeSet
		r        cubeSet
		expected bool
	}

	tests := []testGameSetValidate{
		{
			cs: cubeSet{
				red:   1,
				green: 2,
				blue:  3,
			},
			r: cubeSet{
				red:   1,
				green: 1,
				blue:  1,
			},
			expected: false,
		},
		{
			cs: cubeSet{
				red:   1,
				green: 2,
				blue:  3,
			},
			r: cubeSet{
				red:   1,
				green: 2,
				blue:  3,
			},
			expected: true,
		},
		{
			cs: cubeSet{
				red:   1,
				green: 2,
				blue:  3,
			},
			r: cubeSet{
				red:   5,
				green: 5,
				blue:  5,
			},
			expected: true,
		},
		{
			cs: cubeSet{
				red:   20,
				green: 8,
				blue:  6,
			},
			r: cubeSet{
				red:   12,
				green: 13,
				blue:  14,
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		gotValidation := validateGameSet(tc.cs, tc.r)
		if gotValidation != tc.expected {
			t.Errorf("expected: gameID %t, got %t\n", tc.expected, gotValidation)
		}
	}
}

func TestDay02(t *testing.T) {

	type games struct {
		gs       []string
		r        cubeSet
		expected int
	}

	gameLines := []string{
		"Game 1: 12 blue; 2 green, 13 blue, 19 red; 13 red, 3 green, 14 blue",
		"Game 2: 12 blue, 1 red, 1 green; 1 red, 12 blue, 3 green; 5 green, 1 red, 9 blue; 1 red, 7 blue, 4 green",
		"Game 3: 1 red; 12 blue, 15 red; 1 green, 10 red, 2 blue; 1 green, 3 red, 9 blue",
		"Game 4: 6 blue, 5 green; 2 blue, 6 green, 6 red; 11 blue, 5 red; 6 green, 11 red, 7 blue; 4 green, 10 red; 1 green, 7 red, 13 blue",
		"Game 5: 10 green, 1 red, 2 blue; 3 red, 4 green, 4 blue; 5 green, 5 red",
		"Game 9: 1 blue, 11 red, 9 green; 8 red, 1 blue, 9 green; 4 blue, 16 red, 9 green; 8 green, 3 blue, 6 red; 8 green, 11 red, 3 blue; 11 red, 2 blue",
		"Game 15: 14 blue, 9 green, 1 red; 2 red, 15 blue, 12 green; 1 blue, 2 green, 1 red; 1 red, 16 green, 15 blue; 1 red, 12 green, 8 blue; 1 red, 17 blue",
		"Game 23: 14 red, 2 blue, 9 green; 9 green, 1 blue, 4 red; 9 red, 1 green, 1 blue; 6 green; 3 blue, 1 green, 9 red; 1 blue, 2 red",
		"Game 67: 8 red, 4 blue, 6 green; 4 blue, 8 red, 2 green; 1 green, 6 red, 2 blue; 10 red, 1 green, 2 blue; 1 blue, 5 red; 2 red, 1 green, 2 blue",
		"Game 71: 9 green, 2 blue, 3 red; 5 red; 1 red, 1 blue, 5 green",
		"Game 87: 2 blue, 2 red, 10 green; 8 green, 9 red, 1 blue; 11 red, 1 green, 4 blue; 13 red, 1 blue; 11 green, 16 red, 3 blue",
		"Game 97: 5 green, 13 red, 7 blue; 2 blue, 12 red, 6 green; 10 blue, 11 red, 3 green; 4 green, 11 blue, 15 red; 8 green, 16 blue, 1 red; 15 blue, 4 red, 5 green",
		"Game 100: 8 red, 4 blue, 4 green; 10 blue, 3 red, 4 green; 10 green, 4 red; 18 red, 9 blue, 2 green; 12 red, 4 green, 2 blue",
	}
	restriction := cubeSet{
		red:   12,
		green: 13,
		blue:  14,
	}

	tests := []games{
		{
			gs:       gameLines,
			r:        restriction,
			expected: 2 + 4 + 5 + 67 + 71, // Sum of valid GameIds
		},
	}

	for _, tc := range tests {
		gotSum := day02part1(tc.gs, tc.r)
		if gotSum != tc.expected {
			t.Errorf("expected: sum %d, got %d\n", tc.expected, gotSum)
		}
	}

}
