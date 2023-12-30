// Day 1, part 2.

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse_number(line string, re_fwd regexp.Regexp, re_rev regexp.Regexp) int {
	tens_digit := parse_tens_digit(line, re_fwd)
	units_digit := parse_units_digit(line, re_rev)
	number := 10*tens_digit + units_digit
	return number
}

func parse_tens_digit(line string, re_fwd regexp.Regexp) int {
	match := re_fwd.FindString(line)
	digit, found := parse_digit(match)
	if !found {
		panic(fmt.Sprintf("Couldn't find tens digit in: %s", line))
	}
	return digit
}

func parse_units_digit(line string, re_rev regexp.Regexp) int {
	match := re_rev.FindString(reverse(line))
	digit, found := parse_digit(reverse(match))
	if !found {
		panic(fmt.Sprintf("Couldn't find units digit in: %s", line))
	}
	return digit
}

func parse_digit(digit string) (number int, found bool) {
	found = true
	switch digit {
	case "one", "1":
		number = 1
	case "two", "2":
		number = 2
	case "three", "3":
		number = 3
	case "four", "4":
		number = 4
	case "five", "5":
		number = 5
	case "six", "6":
		number = 6
	case "seven", "7":
		number = 7
	case "eight", "8":
		number = 8
	case "nine", "9":
		number = 9

	default:
		number = -1
		found = false
	}
	return number, found
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	file, err := os.Open("../day01/input_d1.txt")
	defer file.Close()
	check(err)

	scanner := bufio.NewScanner(file)
	sum := 0
	fwd_regexp_str := "one|two|three|four|five|six|seven|eight|nine"
	re_fwd := regexp.MustCompile("([1-9]|" + fwd_regexp_str + ")")
	re_rev := regexp.MustCompile("([1-9]|" + reverse((fwd_regexp_str)) + ")")

	for scanner.Scan() {
		sum += parse_number(scanner.Text(), *re_fwd, *re_rev)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
