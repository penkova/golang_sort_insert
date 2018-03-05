package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello")
}
func Hello(){
	fmt.Println("hello")
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
