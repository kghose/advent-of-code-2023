// Day 2, part 2.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../day02/input.txt")
	defer file.Close()
	check(err)

	scanner := bufio.NewScanner(file)
	power_sum := 0
	for scanner.Scan() {
		power_sum += game_power(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(power_sum)
}

type Trial struct {
	red, green, blue int
}

func game_power(line string) int {
	min_game := Trial{}
	for _, trial_str := range strings.Split(strings.Split(line, ":")[1], ";") {
		trial := parse_trial(trial_str)
		if min_game.red < trial.red {
			min_game.red = trial.red
		}
		if min_game.green < trial.green {
			min_game.green = trial.green
		}
		if min_game.blue < trial.blue {
			min_game.blue = trial.blue
		}
	}
	return min_game.red * min_game.green * min_game.blue
}

func parse_trial(trial_str string) Trial {
	trial := Trial{}
	for _, data := range strings.Split(trial_str, ",") {
		update_trial(&trial, data)
	}
	return trial
}

func update_trial(trial *Trial, data string) {
	data = strings.TrimSpace(data)
	fragments := strings.Split(data, " ")
	count, _ := strconv.Atoi(strings.TrimSpace(fragments[0]))
	switch strings.TrimSpace(fragments[1]) {
	case "red":
		trial.red = count
	case "green":
		trial.green = count
	case "blue":
		trial.blue = count

	default:
		panic(fmt.Sprintf("Unexpected data: %s", data))
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
