// Day 1, part 1.

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("input_d1.txt")
	defer file.Close()
	check(err)

	re := regexp.MustCompile("[0-9]")

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		var (
			line []byte        = scanner.Bytes()
			re   regexp.Regexp = *re
		)
		matches := re.FindAllIndex(line, -1)
		first_digit := line[matches[0][0]]
		second_digit := line[matches[len(matches)-1][0]]
		number, _ := strconv.Atoi(fmt.Sprintf("%c%c", first_digit, second_digit))
		sum += number
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
