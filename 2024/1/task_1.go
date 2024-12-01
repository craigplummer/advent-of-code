package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Counts[S []E, E comparable](s S) map[E]int {
	h := map[E]int{}
	for _, v := range s {
		h[v] = h[v] + 1
	}
	return h
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	location_1 := []int{}
	location_2 := []int{}

	for scanner.Scan() {
		locations := strings.Split(scanner.Text(), "   ")

		location_1_int, _ := strconv.Atoi(locations[0])
		location_2_int, _ := strconv.Atoi(locations[1])

		location_1 = append(location_1, location_1_int)
		location_2 = append(location_2, location_2_int)
	}

	slices.Sort(location_1)
	slices.Sort(location_2)

	total_distance := 0
	similarity := 0
	counts := Counts(location_2)

	for index, location := range location_1 {
		if location_2[index] > location {
			total_distance += location_2[index] - location
		} else {
			total_distance += location - location_2[index]
		}

		similarity += location * counts[location]
	}

	log.Println(total_distance)
	log.Println(similarity)
}
