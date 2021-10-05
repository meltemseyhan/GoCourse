package main

import "testing"

func Test_Basics(t *testing.T) {
	basic()
	arrays()
	maps()
	structs()
	slices()
}

func BenchmarkBasics(t *testing.B) {
	basic()
	arrays()
	maps()
	structs()
	slices()
}
