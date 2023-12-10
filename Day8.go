package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseDay8() (string, map[string]string, map[string]string) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	left := map[string]string{}
	right := map[string]string{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ")", "")

		var h, l, r string
		_, _ = fmt.Sscanf(line, "%s = (%s %s", &h, &l, &r)
		left[h] = l
		right[h] = r
	}

	return instructions, left, right
}

func Day8Part1() {
	instructions, left, right := parseDay8()

	n := len(instructions)
	now := "AAA"
	for count := 1; ; count++ {
		if instructions[(count-1)%n] == 'L' {
			now = left[now]
		} else {
			now = right[now]
		}

		if now == "ZZZ" {
			fmt.Println(count)
			return
		}
	}
}

func Day8Part2() {
	instructions, left, right := parseDay8()

	curr := []string{}
	for k := range left {
		if k[2] == 'A' {
			curr = append(curr, k)
		}
	}

	periods := make([]int, len(curr))

	for i := 1; ; i++ {
		for j, k := range curr {

			switch instructions[(i-1)%len(instructions)] {
			case 'L':
				curr[j] = left[k]
			case 'R':
				curr[j] = right[k]
			}

			if curr[j][2] == 'Z' && periods[j] == 0 {
				periods[j] = i
			}
		}

		if complete(periods) {
			fmt.Println(lcmm(periods))
			return
		}
	}
}

func complete(xs []int) bool {
	for _, x := range xs {
		if x == 0 {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcmm(xs []int) int {
	lcm := func(a, b int) int { return a * b / gcd(a, b) }

	result := 1
	for _, n := range xs {
		result = lcm(result, n)
	}
	return result
}
