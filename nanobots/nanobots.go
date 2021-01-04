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

// frequencyMap creates a map of runes to how many times they appear in the string
func frequencyMap(signal string) map[rune]int {
	frequency := make(map[rune]int)
	for _, char := range signal {
		frequency[char]++
	}
	return frequency
}

// mostFrequent takes a frequency map and determines the most frequently seen rune
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

// frequencyAfterTarget creates a map of runes to how many times they appear in the string
// immediately following the target rune
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
