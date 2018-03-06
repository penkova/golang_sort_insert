package main

import (
	"fmt"
	"sort"
)

func main() {
	//a := []int{2,1,3,6,4,5,8,7,10,9}
	Hello()
	//InsertionSort(a)
}
func Hello(){
	fmt.Println("hello")
}

func InsertionSort(slc []int) {
	for i, key := range slc{
		preIndx := i-1// храним индекс предыдущего элемента
			// пока индекс(preIndx) != 0 и значение предыдущего элемента массива больше текущего
			if preIndx >= 0 && slc[preIndx]>key{
				//перестановка элементов
				slc[preIndx + 1] = slc[preIndx]
				slc[preIndx] = key
				preIndx--
			}
	}
	//fmt.Println(slc)
}

func StandardSort(arr []int) {
	sort.Ints(arr)
}
