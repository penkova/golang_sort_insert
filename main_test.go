package main

import (
	"fmt"
	"testing"
	"time"
)

func TestInsertionSort(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a2 := []int{10, 8, 6, 5, 4, 7, 9, 3, 2, 1}
	t0 := time.Now()
	InsertionSort(a2)
	t1 := time.Now()
	mytime:=t1.Sub(t0)
	fmt.Printf("The call insertion sorted %v ns to run.\n", mytime.Nanoseconds())
	for i, val := range a2 {
		if val != a1[i] {
			t.Error("Exected insertion sort", a1[i], "obtained: ", val)
		}
	}

}

func TestStandardSort(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a2 := []int{10, 8, 6, 5, 4, 7, 9, 3, 2, 1}
	t3 := time.Now()
	StandardSort(a2)
	t4 := time.Now()
	mytime:=t4.Sub(t3)
	fmt.Printf("The call standart sorted %v ns to run.\n", mytime.Nanoseconds())
	for i, val := range a2 {
		if val != a1[i] {
			t.Error("Expected standarted sort:", a1[i], "obtained: ", val)
		}
	}
}
