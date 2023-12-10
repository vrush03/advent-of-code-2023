package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Extraction struct {
	Red, Blue, Green int
}

func Day2Part1Main() {
	db := parseDay2()

	var sum int
	for game, exs := range db {
		if possible(exs) {
			sum += game
		}
	}

	fmt.Println(sum)
}

func Day2Part2Main() {
	db := parseDay2()

	var sum int
	for _, exs := range db {
		m := max(exs)
		sum += m.Blue * m.Red * m.Green
	}

	fmt.Println(sum)
}

func max(exs []Extraction) Extraction {
	var maximum Extraction
	for _, ex := range exs {
		if ex.Red > maximum.Red {
			maximum.Red = ex.Red
		}
		if ex.Green > maximum.Green {
			maximum.Green = ex.Green
		}
		if ex.Blue > maximum.Blue {
			maximum.Blue = ex.Blue
		}
	}
	return maximum
}

func possible(exs []Extraction) bool {
	for _, ex := range exs {
		if ex.Blue > 14 || ex.Red > 12 || ex.Green > 13 {
			return false
		}
	}
	return true
}

func parseDay2() map[int][]Extraction {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var db = map[int][]Extraction{}
	for scanner.Scan() {
		line := scanner.Text()
		game, ex := parseLine(line)
		db[game] = ex
	}
	return db
}

func parseLine(line string) (int, []Extraction) {
	ff := strings.Split(line, ":")
	var game int
	_, _ = fmt.Sscanf(ff[0], "Game %d", &game)

	extractions := []Extraction{}

	for _, ex := range strings.Split(ff[1], ";") {
		var extraction Extraction

		for _, c := range strings.Split(ex, ",") {
			var stones int
			var color string

			_, _ = fmt.Sscanf(c, "%d %s", &stones, &color)

			switch color {
			case "red":
				extraction.Red = stones
			case "blue":
				extraction.Blue = stones
			case "green":
				extraction.Green = stones
			}
		}
		extractions = append(extractions, extraction)
	}

	return game, extractions
}
