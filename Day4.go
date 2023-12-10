package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	ID   int
	Win  map[int]struct{}
	Hand []int
}

const Cards int = 202

func Day4Part1() {
	cards := parseDay4()

	sum := 0
	for _, c := range cards {
		sum += calcPoints(c)
	}
	fmt.Println(sum)
}

func (c Card) Matching() int {
	match := 0
	for _, h := range c.Hand {
		if _, ok := c.Win[h]; ok {
			match++
		}
	}
	return match
}

func Day4Part2() {
	cards := parseDay4()

	// ID -> count
	pocket := map[int]int{}

	for i := 1; i <= Cards; i++ {
		pocket[i] = 1
	}

	for i := 1; i <= Cards; i++ {
		match := cards[i].Matching()
		for j := 1; j <= match; j++ {
			pocket[i+j] += pocket[i]
		}
	}

	sum := 0
	for _, count := range pocket {
		sum += count
	}
	fmt.Println(sum)
}

func calcPoints(c Card) int {
	count := 0
	for _, have := range c.Hand {
		if _, ok := c.Win[have]; ok {
			// fmt.Println(have)
			count++
		}
	}

	// fmt.Println(c)
	// fmt.Println(count)
	if count == 0 {
		return 0
	}
	return int(math.Pow(2, float64(count-1)))
}

func parseDay4() map[int]Card {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	cards := map[int]Card{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		c := Card{
			ID:   0,
			Win:  map[int]struct{}{},
			Hand: []int{},
		}

		temp := strings.Split(line, ":")
		_, _ = fmt.Sscanf(temp[0], "Card %d", &c.ID)

		temp2 := strings.Split(temp[1], "|")

		for _, num := range strings.Fields(temp2[0]) {
			number, _ := strconv.Atoi(num)
			c.Win[number] = struct{}{}
		}

		for _, num := range strings.Fields(temp2[1]) {
			number, _ := strconv.Atoi(num)
			c.Hand = append(c.Hand, number)
		}
		cards[c.ID] = c
	}

	return cards
}
