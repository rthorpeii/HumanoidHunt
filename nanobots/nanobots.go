package nanobots

import "reaktor/input"

var puzzleInput = "./nanobots/input.txt"

// Solve solves the puzzle
func Solve() string {
	rawInput := input.ReadInput(puzzleInput)
	mostFrequentChar := mostFrequent(frequencyMap(rawInput))

	password := ""
	for string(mostFrequentChar) != ";" {
		password += string(mostFrequentChar)
		frequencyFromCurrent := frequencyAfterTarget(rawInput, mostFrequentChar)
		mostFrequentChar = mostFrequent(frequencyFromCurrent)
	}
	return password
}

func mostFrequent(frequency map[rune]int) rune {
	var maxChar rune
	maxFrequency := 0

	for char, value := range frequency {
		if value > maxFrequency {
			maxChar = char
			maxFrequency = value
		}
	}
	return maxChar
}

func frequencyMap(signal string) map[rune]int {
	frequency := make(map[rune]int)
	for _, char := range signal {
		frequency[char]++
	}
	return frequency
}

func frequencyAfterTarget(signal string, target rune) map[rune]int {
	frequency := make(map[rune]int)
	for i := 0; i < len(signal)-1; i++ {
		char := rune(signal[i])
		if char == target {
			frequency[rune(signal[i+1])]++
		}
	}
	return frequency
}
