package main

import (
	"fmt"
	"reaktor/android"
	"reaktor/nanobots"
	"reaktor/tattoo"
)

func main() {
	fmt.Println("Part 1 - Tatoo: ", tattoo.Solve())
	fmt.Println("Part 2 - Nanobots: ", nanobots.Solve())
	fmt.Println("Part 3 - Android: ", android.Solve())
}
