package main

import "testing"

func TestPart1(t *testing.T) {
	input := getInput("input_test")
	got := part1(input)
	want := 24000

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := getInput("input_test")
	got := part2(input)
	want := 45000

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
