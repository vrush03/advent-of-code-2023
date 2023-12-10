package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Day6Part2() {
	start := time.Now()
	raceTime, distance := parseB()
	// mul := 1
	count := 0
	for t := 0; t <= raceTime; t++ {
		if (raceTime-t)*t > distance {
			count++
		}
	}
	fmt.Println(count)
	fmt.Println(time.Since(start))
	// }	for index, time := range times {
	// 	count := 0
	// 	for t := 0; t <= time; t++ {
	// 		if (time-t)*t > distances[index] {
	// 			count++
	// 		}
	// 	}
	// 	mul = mul * count
	// }
	//
	// fmt.Println(mul)
}

func Day6Part1() {
	times, distances := parseA()
	mul := 1
	for index, time := range times {
		count := 0
		for t := 0; t <= time; t++ {
			if (time-t)*t > distances[index] {
				count++
			}
		}
		mul = mul * count
	}

	fmt.Println(mul)
}

func parseA() ([]int, []int) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	//
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var times []int
	for _, time := range strings.Fields(scanner.Text())[1:] {
		n, _ := strconv.Atoi(time)
		times = append(times, n)
	}

	scanner.Scan()
	var distances []int
	for _, distance := range strings.Fields(scanner.Text())[1:] {
		n, _ := strconv.Atoi(distance)
		distances = append(distances, n)
	}

	return times, distances

}

func parseB() (int, int) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	//
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	raceTime, _ := strconv.Atoi(strings.Join(strings.Fields(scanner.Text())[1:], ""))
	scanner.Scan()
	raceDist, _ := strconv.Atoi(strings.Join(strings.Fields(scanner.Text())[1:], ""))

	return raceTime, raceDist

}
