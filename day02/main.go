// Day 2, part 1.
// Assumes games are listed and numbered in order and all lines have valid data

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)

	scanner := bufio.NewScanner(file)
	sum_id := 0
	game_no := 1
	for scanner.Scan() {
		if valid_game(scanner.Text()) {
			sum_id += game_no
		}
		game_no++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum_id)
}

func valid_game(line string) bool {
	data := strings.Split(strings.Split(line, ":")[1], ";")
	for _, datum := range data {
		cube_data := strings.Split(datum, ",")
		for _, cube_datum := range cube_data {
			if !valid_count(cube_datum) {
				return false
			}
		}

	}
	return true
}

func valid_count(cube_datum string) bool {
	datum := strings.TrimSpace(cube_datum)
	fragments := strings.Split(datum, " ")
	count, _ := strconv.Atoi(fragments[0])
	switch fragments[1] {
	case "red":
		return count <= 12
	case "green":
		return count <= 13
	case "blue":
		return count <= 14
	}
	panic(fmt.Sprintf("Unexpected data: %s", datum))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
