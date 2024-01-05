/* Day 3, part 1

We assume all lines are the same length.

Keep three raw lines L-1, L, L+1, where L is the current line.
Keep a symbol adjacency buffer L' that stores a marker 'X' to
indicate which columns are symbol adjacent. If a number in L falls
on an 'X' that is a part number.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	part_number_sum := 0
	for schematic := Schematic(file); schematic.has_line_below; {
		part_number_sum += schematic.read_and_process_line()
	}
	fmt.Println(part_number_sum)
}

type SchematicFragment struct {
	line_above,
	this_line,
	line_below []byte
	line_len       int
	scanner        *bufio.Scanner
	has_line_below bool
}

func Schematic(file *os.File) SchematicFragment {
	scanner := bufio.NewScanner(file)
	_line_below, has_line_below := read_line_below(scanner, 0)
	line_below := empty_line(len(_line_below))
	copy(line_below, _line_below)
	return SchematicFragment{
		line_above:     empty_line(len(line_below)),
		this_line:      empty_line(len(line_below)),
		line_below:     line_below,
		line_len:       len(line_below),
		scanner:        scanner,
		has_line_below: has_line_below,
	}
}

func empty_line(l int) []byte {
	line := make([]byte, l)
	for i := 0; i < l; i++ {
		line[i] = '.'
	}
	return line
}

func read_line_below(scanner *bufio.Scanner, line_len int) (line_below []byte, has_line_below bool) {
	has_line_below = scanner.Scan()
	if has_line_below {
		line_below = scanner.Bytes()
	} else {
		line_below = empty_line(line_len)
	}
	return line_below, has_line_below
}

func (f *SchematicFragment) read_and_process_line() int {
	line_below, has_line_below := read_line_below(f.scanner, f.line_len)
	f.load_next(line_below)
	f.has_line_below = has_line_below
	return f.process_line()
}

func (f *SchematicFragment) load_next(line []byte) {
	copy(f.line_above, f.this_line)
	copy(f.this_line, f.line_below)
	copy(f.line_below, line)
}

func (f *SchematicFragment) process_line() int {
	part_number_sum := 0

	symbol_adjacency := f.get_symbol_adjacency()

	we_are_parsing_digits := false
	digits := []byte{}
	is_part_number := false
	for i := 0; i < len(f.this_line); i++ {
		if is_digit(f.this_line[i]) {
			digits = append(digits, f.this_line[i])
			we_are_parsing_digits = true
		} else {
			if we_are_parsing_digits && is_part_number {
				number, _ := strconv.Atoi(string(digits))
				part_number_sum += number
			}
			we_are_parsing_digits = false
			is_part_number = false
			digits = []byte{}
		}

		if we_are_parsing_digits && symbol_adjacency[i] == 'X' {
			is_part_number = true
		}
	}

	if we_are_parsing_digits && is_part_number {
		number, _ := strconv.Atoi(string(digits))
		part_number_sum += number
	}

	return part_number_sum
}

func (f *SchematicFragment) get_symbol_adjacency() []byte {
	symbol_adjacency := empty_line(len(f.this_line))

	for i := 0; i < len(f.this_line); i++ {
		if has_symbol(f.line_above[i]) ||
			has_symbol(f.this_line[i]) ||
			has_symbol(f.line_below[i]) {
			if i > 0 {
				symbol_adjacency[i-1] = 'X'
			}
			symbol_adjacency[i] = 'X'
			if i < len(f.this_line)-1 {
				symbol_adjacency[i+1] = 'X'
			}
		}
	}
	return symbol_adjacency
}

func is_digit(b byte) bool {
	if ('0' <= b) && (b <= '9') {
		return true
	} else {
		return false
	}
}

func has_symbol(b byte) bool {

	switch {
	case b == '.':
		return false
	case b < '0':
		return true
	case b > '9':
		return true
	}
	return false
}
