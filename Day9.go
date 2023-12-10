package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day9Part2Main() {
	list := parseDay9()

	sum := 0
	for _, row := range list {
		p := buildPascal(row)

		p[len(p)-1] = append([]int{0}, p[len(p)-1]...)

		for i := len(p) - 2; i >= 0; i-- {
			p[i] = append([]int{p[i][0] - p[i+1][0]}, p[i]...)
		}

		sum += p[0][0]
	}

	fmt.Println(sum)
}

func Day9Part1Main() {
	list := parseDay9()

	sum := 0
	for _, row := range list {
		p := buildPascal(row)

		p[len(p)-1] = append(p[len(p)-1], 0)

		for i := len(p) - 2; i >= 0; i-- {
			p[i] = append(p[i], p[i][len(p[i])-1]+p[i+1][len(p[i+1])-1])
		}

		sum += p[0][len(p[0])-1]
	}

	fmt.Println(sum)
}

func parseDay9() [][]int {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	list := [][]int{}
	for scanner.Scan() {
		sequence := []int{}
		for _, f := range strings.Fields(scanner.Text()) {
			n, _ := strconv.Atoi(f)
			sequence = append(sequence, n)
		}
		list = append(list, sequence)
	}
	return list
}

func buildPascal(firstSeq []int) [][]int {
	p := [][]int{firstSeq}
	for {
		last := len(p) - 1

		done := true
		newRow := make([]int, len(p[last])-1)
		for i := 0; i < len(p[last])-1; i++ {
			newRow[i] = p[last][i+1] - p[last][i]
			if newRow[i] != 0 {
				done = false
			}
		}
		p = append(p, newRow)

		if done {
			return p
		}
	}
}
