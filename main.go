package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	fmt.Println("Hello")
	hours, _ := time.ParseDuration("10h")
	complex, _ := time.ParseDuration("1h10m10s")

	fmt.Println(hours)
	fmt.Println(complex)
	fmt.Printf("there are %.0f seconds in %v\n", complex.Seconds(), complex)
}

func InsertionSort(arr []int) {
	//left to right
	for i := 0; i < len(arr); i++ {
		key := arr[i]
		tmp := i
		//right to left
		for j := i - 1; j > -1; j-- {
			if arr[j] < key {
				break
			}
			arr[j+1] = arr[j]
			tmp = j
		}
		arr[tmp] = key
	}
}

func StandardSort(arr []int) {
	sort.Ints(arr)
}
