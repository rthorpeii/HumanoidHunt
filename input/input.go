package input

import (
	"io/ioutil"
	"log"
	"strings"
)

// ReadInput takes an input file and converts it into a string
func ReadInput(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Failed to convert input to string", err)
	}
	return string(data)
}

// Slice parses the input into a slice
func Slice(file string) []string {
	return strings.Split(ReadInput(file), "\n")
}
