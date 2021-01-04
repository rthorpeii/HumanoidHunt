package tattoo

import (
	"log"
	"reaktor/input"
	"strconv"
)

var testInput = "./tattoo/input.txt"

// Solve the puzzle
func Solve() string {
	rawInput := input.Slice(testInput)
	password := ""
	for _, value := range rawInput {
		password += string(findInvalid(value))
	}
	return password
}

// findInvalid returns the character represented by the first invalid byte
// following a sequence of valid bytes in the stream
func findInvalid(stream string) rune {
	numBytes := len(stream) / 8
	byteOffset := 0
	valid := false
	for byteOffset < numBytes {
		byteIndex := byteOffset * 8
		currentByte := stream[byteIndex : byteIndex+8]
		byteValue, _ := strconv.ParseInt(currentByte, 2, 8)
		intByteValue := int(byteValue)
		if intByteValue < numBytes {
			valid = true
			byteOffset = intByteValue
			continue
		} else if valid {
			return rune(intByteValue)
		}
		byteOffset++
	}
	log.Fatal("No character found")
	return rune(0)
}
