package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Map []Transform

type Transform struct {
	Destination, Source, Length int
}

func Day5Part2() {
	seeds, maps := parseDay5()

	for loc := 0; ; loc++ {
		x := loc
		for m := len(maps) - 1; m >= 0; m-- {
			x = maps[m].InverseConvert(x)
		}

		for s := 0; s < len(seeds); s += 2 {
			if x >= seeds[s] && x <= seeds[s]+seeds[s+1] {
				fmt.Println(loc)
				return
			}
		}
	}
}

func Day5Part1() {
	seeds, maps := parseDay5()

	minLocation := 1<<31 - 1
	for _, s := range seeds {
		for _, m := range maps {
			s = m.Convert(s)
		}

		if s < minLocation {
			minLocation = s
		}
	}

	fmt.Println(minLocation)
}

func (m Map) Convert(from int) int {
	for _, t := range m {
		if from >= t.Source && from < t.Source+t.Length {
			return t.Destination + (from - t.Source)
		}
	}
	return from
}

func parseDay5() ([]int, [7]Map) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var seeds []int
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	for i := 1; i < len(fields); i++ {
		num, _ := strconv.Atoi(fields[i])
		seeds = append(seeds, num)
	}

	return seeds, parseMaps(scanner)
}

func (m Map) InverseConvert(to int) int {
	for _, t := range m {
		if to >= t.Destination && to < t.Destination+t.Length {
			return t.Source + (to - t.Destination)
		}
	}
	return to
}

func parseMaps(scanner *bufio.Scanner) [7]Map {
	maps := [7]Map{}
	t := -1
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.Contains(line, "map:") {
			t++
			continue
		}

		fields := strings.Fields(line)

		dest, _ := strconv.Atoi(fields[0])
		source, _ := strconv.Atoi(fields[1])
		length, _ := strconv.Atoi(fields[2])

		maps[t] = append(maps[t], Transform{
			Destination: dest,
			Source:      source,
			Length:      length,
		})
	}
	return maps
}
