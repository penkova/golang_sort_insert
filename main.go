package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{7, 3, 4, 1, 5, 6, 2}
	fmt.Println("None norting slice", a)
	//quicksort(a)
	MyQuickSort(a)
	fmt.Println("Sorting slice", a)
}

func Swap(slc []int, a, b int) []int {
	slc[a], slc[b] = slc[b], slc[a]
	return slc
}

func MyQuickSort(slc []int) {
	left := 0
	right := len(slc) - 1
	//qsort(slc,left, right)
	//myQuickSort(slc,left, right)
	myQuickSort1(slc, left, right)
	//myQuickSort2(slc, left, right)
	//myQuickSort3(slc,left, right)
}

//работает! лучший варик!
func quicksort(list []int) {
	l := len(list)
	if l <= 1 {
		return
	}
	b := 0
	e := l - 1
	key := list[e/2]
	fmt.Println("center: ", key, "slice: ", list)
	for {
		for list[b] < key {
			b++
		}
		for list[e] > key {
			e--
		}
		if b >= e {
			break
		}
		Swap(list, b, e)
		b++
		e--
	}

	quicksort(list[:b])
	quicksort(list[e+1:])
}

//работает!
func qsort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	x := arr[i]
	for i < j {
		for i < j && arr[j] >= x {
			j--
		}
		if i < j {
			arr[i] = arr[j]
		}
		for i < j && (arr[i] <= x) {
			i++
		}
		if i < j {
			arr[j] = arr[i]
		}

	}
	arr[i] = x
	qsort(arr, l, i-1)
	qsort(arr, i+1, r)
}

//не работает, но есть идея как доработать
func myQuickSort(slc []int, left, right int) {
	if left >= right {
		return
	}
	center := slc[left+right/2]
	i := left
	j := right

	for i < j {
		for slc[i] < center {
			i++
		}
		for slc[j] > center {
			j--
		}
		if i <= j {
			slc = Swap(slc, i, j)
			i++
			j--
		}
	}

	//
	//myQuickSort(slc, left, i-1)
	//myQuickSort(slc, i+1, right)

}



func myQuickSort1(slc []int, left, right int) {
	l := left
	r := right
	center:= l+r/2
	if slc[center] < slc[right]{
		tmp :=slc[center]
		slc[center] = slc[left]
		slc[left] = tmp
		left++
		center++
	}
	if slc[center] > slc[right]{
		tmp :=slc[center]
		slc[center] = slc[right-1]
		slc[right-1] = tmp
		right--
	}
	MyQuickSort(slc[:left])
	MyQuickSort(slc[right:])
}

func myQuickSort2(slc []int, left, right int) {
	//if left < right {
	//	key := slc[left]
	//	low := left
	//	high := right
	//	for low < high {
	//		for low < high && slc[high] > key {
	//			high--
	//		}
	//		slc[low] = slc[high]
	//		for low < high && slc[low] < key {
	//			low++
	//		}
	//		slc[high] = slc[low]
	//	}
	//	slc[low] = key
	//
	if left >= right {
		return //slice finished
	}
	l := left
	r := right
	//индекс и значение нашего опорного(центральный центральный) элемента
	centerId := right / 2
	centerVal := slc[centerId]
	//fmt.Println("First index", l, "  Last index", r, "   Central index", centerId, "| Value:", centerVal)
	fmt.Println(" Central value", centerVal, "| Slice:", slc)
	// пока левая и права граница не сойдутся
	for l < r {
		// значения больше center
		for l < r && slc[r] > centerVal {
			r--
		}
		slc[l]=slc[r]
		// пока значения меньше center
		for  l < r && slc[l] < centerVal {
			l++
		}
		slc[r]=slc[l]
	}
	slc[l]=centerVal
	myQuickSort2(slc, left, l-1)
	myQuickSort2(slc, l+1, right)
}


//работает неправильно
func myQuickSort3(slc []int, left, right int) {
	if left >= right {
		fmt.Println("Woow, slice ------> 0")
		return
	}
	l := left
	r := right
	//индекс и значение нашего опорного(центральный центральный) элемента
	centerId := right / 2
	centerVal := slc[centerId]
	fmt.Println("First index", l, "  Last index", r, "   Central index", centerId, "| Value:", centerVal)
	// пока левая и права граница не сойдутся
	for l <= r {
		// значения больше center
		for slc[r] > centerVal {
			r--
			fmt.Println(" >central", slc[r])
		}
		// пока значения меньше center
		for slc[l] < centerVal {
			l++
			fmt.Println(" <central", slc[l])
		}

		if l <= r {
			// меняем местами значение c индексом i на месте c индексом j
			slc[l], slc[r] = slc[r], slc[l]
			l++
			r--
		}
	}
	// рекурсивно вызываем сортировку...
	if r > left {
		myQuickSort3(slc, left, r)
	}
	if l < right {
		myQuickSort3(slc, l, right)
	}
	return
}

func Hello() {
	fmt.Println("hello")
}

func InsertionSort(slc []int) {
	for i, key := range slc {
		preIndx := i - 1 // храним индекс предыдущего элемента
		// пока индекс(preIndx) != 0 и значение предыдущего элемента массива больше текущего
		if preIndx >= 0 && slc[preIndx] > key {
			//перестановка элементов
			slc[preIndx+1] = slc[preIndx]
			slc[preIndx] = key
			preIndx--
		}
	}
	//fmt.Println(slc)
}

func StandardSort(arr []int) {
	sort.Ints(arr)
}
