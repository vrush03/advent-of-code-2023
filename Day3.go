package main

import (
	"bufio"
	"fmt"
	"os"
)

type C struct{ i, j int }

const Size = 140

func Day3Part1Main() {
	m := parseDay3()

	sum := 0
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {

			c := C{i, j}
			ch := m[c]

			if isDigit(ch) {
				valid := false
				num := 0
				for isDigit(ch) {
					num = 10*num + int(ch-'0')
					valid = valid || isValid(m, c)
					j++
					c = C{i, j}
					ch = m[c]
				}

				if valid {
					sum += num
				}
			}
		}
	}
	fmt.Println(sum)
}

func isValid(m map[C]byte, c C) bool {
	delta := []int{-1, 0, 1}
	for _, di := range delta {
		for _, dj := range delta {
			if di|dj != 0 {
				ch, exists := m[C{c.i + di, c.j + dj}]
				if exists && !isDigit(ch) {
					return true
				}
			}
		}
	}
	return false
}

func Day3Part2Main() {
	m := parseDay3()

	sum := 0
	for i := 0; i < 140; i++ {
		for j := 0; j < 140; j++ {
			c := C{i, j}
			ch := m[c]

			if ch == '*' {
				sum += gearRatio(m, c)
			}
		}
	}
	fmt.Println(sum)
}

func (c C) Left() C   { return C{c.i, c.j - 1} }
func (c C) Right() C  { return C{c.i, c.j + 1} }
func (c C) Top() C    { return C{c.i - 1, c.j} }
func (c C) Bottom() C { return C{c.i + 1, c.j} }

func gearRatio(m map[C]byte, c C) int {

	numbers := []int{}

	if isDigit(m[c.Left()]) {
		numbers = append(numbers, expandNumber(m, c.Left()))
	}

	if isDigit(m[c.Right()]) {
		numbers = append(numbers, expandNumber(m, c.Right()))
	}

	if isDigit(m[c.Top()]) {
		numbers = append(numbers, expandNumber(m, c.Top()))
	} else {
		if isDigit(m[c.Top().Left()]) {
			numbers = append(numbers, expandNumber(m, c.Top().Left()))
		}
		if isDigit(m[c.Top().Right()]) {
			numbers = append(numbers, expandNumber(m, c.Top().Right()))
		}
	}

	if isDigit(m[c.Bottom()]) {
		numbers = append(numbers, expandNumber(m, c.Bottom()))
	} else {
		if isDigit(m[c.Bottom().Left()]) {
			numbers = append(numbers, expandNumber(m, c.Bottom().Left()))
		}
		if isDigit(m[c.Bottom().Right()]) {
			numbers = append(numbers, expandNumber(m, c.Bottom().Right()))
		}
	}

	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	}

	return 0
}

func expandNumber(m map[C]byte, c C) int {
	for isDigit(m[c.Left()]) {
		c.j--
	}
	num := 0
	for ; isDigit(m[C{c.i, c.j}]); c.j++ {
		num = 10*num + int(m[c]-'0')
	}
	return num
}

func parseDay3() map[C]byte {
	// scanner := bufio.NewScanner(os.Stdin)
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := map[C]byte{}

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		for j := 0; j < len(line); j++ {
			c := line[j]
			if c == '.' {
				continue
			}
			m[C{i, j}] = c
		}
	}
	return m
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
