package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{7, 3, 4, 1, 5, 6, 2, 8}
	fmt.Println("Not sorting slice", a)
	//quicksort1(a)
	//MyQuickSort(a)
	InsertionSort(a)
	fmt.Println("Sorting slice", a)
}

func MyQuickSort(slc []int) {
	left := 0
	right := len(slc)
	//qsort(slc, left, right)
	myQuickSort(slc, left, right)
	return
}

//работает!
func myQuickSort(slc []int, left, right int) {
	if left >= right {
		return //slice finished
	}
	l := 0
	r := right - 1
	//индекс и значение нашего опорного(центральный центральный) элемента
	centerId := right / 2
	//fmt.Println("start:", l, "end:", r, "center:", centerId)
	//fmt.Println(" 	Central value", slc[centerId], "| Slice:", slc)
	// пока левая и права граница не сойдутся
	for {
		//fmt.Println("		значение центра", slc[centerId], "| Slice:", slc)
		// по сути движение начинается одновременно с начала и конца до нашего опрного элемента
		for slc[l] < slc[centerId] {
			//fmt.Println("< key center: ", slc[centerId], "slice: ", slc)
			l++
		}
		for slc[r] > slc[centerId] {
			r--
		}
		if l >= r {
			break
		}
		//если же элемент не соответствует условии в FORах, то происходит замена элементов с опорным элементом
		slc[l], slc[r] = slc[r], slc[l]
		l++
		r--
	}
	myQuickSort(slc, left, r)
	myQuickSort(slc, l+1, right-1)
	// OR
	//MyQuickSort(slc[:r])
	//MyQuickSort(slc[r+1:])
	return
}

//func InsertionSort(slc []int) {
//	for i, key := range slc {
//		preIndx := i -1 // храним индекс предыдущего элемента
//		// пока индекс(preIndx) != 0 и значение предыдущего элемента массива больше текущего
//		if preIndx >= 0 && slc[preIndx] > key {
//			slc[preIndx],slc[preIndx+1] = slc[preIndx+1],slc[preIndx]
//			preIndx--
//		}
//	}
//	fmt.Println(slc)
//	return
//}

func InsertionSort(slc []int) {
	for i, key := range slc{
			j := i - 1
			for j >= 0 && slc[j] > key {
				slc[j+1] = slc[j]
				j = j- 1
			}

			slc[j+1] = key
	}
}

func StandardSort(arr []int) {
	sort.Ints(arr)
	return
}






//работает!
func quicksort1(slc []int) {
	l := len(slc)
	if l <= 1 {
		return
	}
	i := 0
	j := l - 1
	key := slc[j/2]
	fmt.Println("center: ", key, "slice: ", slc)
	for {
		for slc[i] < key {
			fmt.Println("< key center: ", key, "slice: ", slc)
			i++
		}
		for slc[j] > key {
			fmt.Println("> kye center: ", key, "slice: ", slc)
			j--
		}
		if i >= j {
			break
		}
		slc[i], slc[j] = slc[j], slc[i]
		i++
		j--
	}
	quicksort1(slc[:j])
	quicksort1(slc[j+1:])
}

//работает!
func quicksort2(arr []int, l int, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	x := arr[i] //опорный элемент выбираем первый
	fmt.Println("pivot: ", x, "slice: ", arr)
	for i < j {
		for i < j && arr[j] >= x {
			fmt.Println("> pivot: ", x, "slice: ", arr)
			j--
		}
		if i < j {
			arr[i] = arr[j]
		}
		for i < j && (arr[i] <= x) {
			fmt.Println("< pivot: ", x, "slice: ", arr)
			i++
		}
		if i < j {
			arr[j] = arr[i]
		}

	}
	arr[i] = x
	quicksort2(arr, l, i-1)
	quicksort2(arr, i+1, r)
}