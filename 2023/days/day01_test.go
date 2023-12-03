package day01

import (
	"testing"
)

func TestDay01(t *testing.T) {

	type test struct {
		input    []string
		expected int
	}

	tests := []test{
		{input: []string{"snqhqmffonettwofourgdkjmbjvjpgxxxpzkm8zfpfcgj"}, expected: 18},
		{input: []string{"9jdpnzgqrf"}, expected: 99},
		{input: []string{"pgblvrqlnjfdtvngfbzpl5njsmvshn5tc"}, expected: 55},
		{input: []string{"nine6one9jnqf1"}, expected: 91},
		{input: []string{"onetwo8kbxqgvsevenmrhqndt"}, expected: 17},
		{input: []string{"bqzpzbzkbs7nprbdmbqseven8kzr1pflnine"}, expected: 79},
		{input: []string{"nine5five2375lhphjk"}, expected: 95},
		{input: []string{"4onetwotrnqlgxgtxxrgxpgsevenddjfd"}, expected: 47},
		{input: []string{"ninezrvbf717six"}, expected: 96},
		{input: []string{"foursdmljtklzldsevenvbqpthree917"}, expected: 47},
		{input: []string{"hkxrxtdjzdzqnrzxfzsix3three"}, expected: 63},
		{input: []string{
			"snqhqmffonettwofourgdkjmbjvjpgxxxpzkm8zfpfcgj",
			"9jdpnzgqrf",
			"pgblvrqlnjfdtvngfbzpl5njsmvshn5tc",
			"nine6one9jnqf1",
			"onetwo8kbxqgvsevenmrhqndt",
			"bqzpzbzkbs7nprbdmbqseven8kzr1pflnine",
			"nine5five2375lhphjk",
			"4onetwotrnqlgxgtxxrgxpgsevenddjfd",
			"ninezrvbf717six",
			"foursdmljtklzldsevenvbqpthree917",
			"hkxrxtdjzdzqnrzxfzsix3three",
		}, expected: 18 + 99 + 55 + 91 + 17 + 79 + 95 + 47 + 96 + 47 + 63,
		},
	}

	for _, tc := range tests {
		got := day01part2(tc.input)
		if got != tc.expected {
			t.Errorf("expected: %d, got %d", tc.expected, got)
		}
	}
}
