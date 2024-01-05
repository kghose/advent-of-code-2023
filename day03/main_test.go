package main

import (
	"bytes"
	"os"
	"testing"
)

func TestSymbolAdjacency(t *testing.T) {

	schematic := SchematicFragment{
		line_above: []byte("..$....."),
		this_line:  []byte("23.....*"),
		line_below: []byte("...*....")}
	//                      .XXXX.XX
	got := schematic.get_symbol_adjacency()
	if !bytes.Equal(got, []byte(".XXXX.XX")) {
		t.Errorf("Symbol adjacency should be '.XXXX.XX', got '%s'", got)
	}
}

func TestPartNumberSum(t *testing.T) {
	schematic := SchematicFragment{
		line_above: []byte("..$....."),
		this_line:  []byte("23...8.*"),
		line_below: []byte("...*....")}
	got := schematic.process_line()
	if got != 23 {
		t.Errorf("Part number sum should be 23, got %d", got)
	}
}

func TestPartNumberSum2(t *testing.T) {
	schematic := SchematicFragment{
		line_above: []byte("........"),
		this_line:  []byte("23..*8.*"),
		line_below: []byte("...*....")}
	got := schematic.process_line()
	if got != 8 {
		t.Errorf("Part number sum should be 8, got %d", got)
	}
}

func TestPartNumberSum3(t *testing.T) {
	schematic := SchematicFragment{
		line_above: []byte(".&......"),
		this_line:  []byte("23..*8.*"),
		line_below: []byte("...*....")}
	got := schematic.process_line()
	if got != 31 {
		t.Errorf("Part number sum should be 31, got %d", got)
	}
}

func TestPartNumberSumEdge1(t *testing.T) {
	schematic := SchematicFragment{
		line_above: []byte("........"),
		this_line:  []byte("......*8"),
		line_below: []byte("........")}
	got := schematic.process_line()
	if got != 8 {
		t.Errorf("Part number sum should be 8, got %d", got)
	}
}

func TestPartNumberSumEdge2(t *testing.T) {
	schematic := SchematicFragment{
		line_above: []byte("........"),
		this_line:  []byte("......8*"),
		line_below: []byte("........")}
	got := schematic.process_line()
	if got != 8 {
		t.Errorf("Part number sum should be 8, got %d", got)
	}
}

func TestSchematicScrolling(t *testing.T) {
	schematic := SchematicFragment{
		line_above: []byte(".&......"),
		this_line:  []byte("23..*8.*"),
		line_below: []byte("..456...")}
	schematic.load_next([]byte(".&......"))
	got := schematic.process_line()
	if got != 456 {
		t.Errorf("Part number sum should be 456, got %d", got)
	}
}

func TestEndToEnd(t *testing.T) {
	file, err := os.Open("test_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	part_number_sum := 0
	for schematic := Schematic(file); schematic.has_line_below; {
		part_number_sum += schematic.read_and_process_line()
	}

	if part_number_sum != 4361 {
		t.Errorf("Part number sum should be 4361, got %d", part_number_sum)
	}
}
