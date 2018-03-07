package main

import (
	"testing"
)

func BenchmarkInsertionSort(b *testing.B) {
	a2 := []int{2, 1, 3, 6, 4, 5, 8, 7, 10, 9}
	for n := 0; n < b.N; n++ {
		InsertionSort(a2)
	}
}

func BenchmarkMyQuickSort(b *testing.B) {
	a2 := []int{2, 1, 3, 6, 4, 5, 8, 7, 10, 9}
	for n := 0; n < b.N; n++ {
		MyQuickSort(a2)
	}
}

func BenchmarkStandardSort(b *testing.B) {
	a2 := []int{2, 1, 3, 6, 4, 5, 8, 7, 10, 9}
	for n := 0; n < b.N; n++ {
		StandardSort(a2)
	}
}

func TestInsertionSort(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	InsertionSort(a2)
	for i, val := range a2 {
		if val != a1[i] {
			t.Error("Exected insertion sort", a1[i], "obtained: ", val)
		}
	}
}

func TestStandardSort(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	StandardSort(a2)
	for i, val := range a2 {
		if val != a1[i] {
			t.Error("Expected standarted sort:", a1[i], "obtained: ", val)
		}
	}
}
func TestMyQuickSort(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	MyQuickSort(a2)
	for i, val := range a2 {
		if val != a1[i] {
			t.Error("Expected standarted sort:", a1[i], "obtained: ", val)
		}
	}
}
