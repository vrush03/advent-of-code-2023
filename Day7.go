package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Game struct {
	Hand string
	Type int
	Bid  int
}

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var valuePart1 = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var valuePart2 = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1, // Joker
}

func Type2(hand string) int {
	count := map[rune]int{}

	for _, c := range hand {
		count[c]++
	}

	assignJokers(count)

	return getTypeByCount(count)
}

func assignJokers(count map[rune]int) {
	for count['J'] > 0 {
		vMaxNotJoker := -1
		var card rune
		for k, v := range count {
			if k == 'J' {
				continue
			}
			if v > vMaxNotJoker {
				vMaxNotJoker = v
				card = k
			}
		}
		count['J']--
		count[card]++
	}
}

func Day7Main() {
	games := parseDay7()

	for i := 0; i < len(games); i++ {
		games[i].Type = Type2(games[i].Hand)
	}

	sort.Slice(games, func(i, j int) bool {
		if games[i].Type < games[j].Type {
			return true
		}
		if games[i].Type > games[j].Type {
			return false
		}

		for c := 0; c < 5; c++ {
			card1, card2 := valuePart2[games[i].Hand[c]], valuePart2[games[j].Hand[c]]
			if card1 < card2 {
				return true
			}
			if card1 > card2 {
				return false
			}
		}
		return false
	})

	winning := 0
	for i, g := range games {
		winning += g.Bid * (i + 1)
	}

	fmt.Println(winning)
}

func parseDay7() []Game {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var games []Game
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		hand := fields[0]
		bid, _ := strconv.Atoi(fields[1])

		games = append(games, Game{
			Hand: hand,
			Bid:  bid,
		})
	}
	return games
}

func getTypeByCount(count map[rune]int) int {
	values := sort.IntSlice{}
	for _, v := range count {
		values = append(values, v)
	}
	sort.Sort(sort.Reverse(values))

	switch values[0] {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if values[1] == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	case 2:
		if values[1] == 2 {
			return TwoPair
		}
		return OnePair
	case 1:
		return HighCard
	}
	log.Fatal("impossible")
	return 0
}

func Type(hand string) int {
	count := map[rune]int{}

	for _, c := range hand {
		count[c]++
	}

	return getTypeByCount(count)
}
