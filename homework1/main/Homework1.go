package main

import "fmt"

func SortSlice(slice []int) {
	if len(slice) <= 1 {
		return
	}

	mid := len(slice) / 2
	left := make([]int, mid)
	right := make([]int, len(slice)-mid)

	copy(left, slice[:mid])
	copy(right, slice[mid:])

	SortSlice(left)
	SortSlice(right)

	merge(slice, left, right)
}

func merge(slice, left, right []int) {
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		slice[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		slice[k] = right[j]
		j++
		k++
	}
}

func IncrementOdd(slice []int) {
	for i := 1; i < len(slice); i += 2 {
		slice[i]++
	}
}

func PrintSlice(slice []int) {
	fmt.Print("[")
	for i, value := range slice {
		fmt.Print(value)
		if i < len(slice)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func ReverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

type sliceFunc func([]int)

func appendFunc(dst sliceFunc, src ...sliceFunc) sliceFunc {
	return func(slice []int) {
		dst(slice)
		for _, fn := range src {
			fn(slice)
		}
	}
}

func main() {
	fmt.Println()
	arr := []int{5, 2, 9, 1, 5, 6}
	SortSlice(arr)
	fmt.Println(arr)
	fmt.Println("--------------------------------------")

	IncrementOdd(arr)
	fmt.Println(arr)
	fmt.Println("--------------------------------------")

	PrintSlice(arr)
	fmt.Println("--------------------------------------")

	ReverseSlice(arr)
	fmt.Println(arr)
	fmt.Println("--------------------------------------")

	printSlice := func(slice []int) {
		fmt.Println(slice)
	}

	doubleValues := func(slice []int) {
		for i := range slice {
			slice[i] *= 2
		}
	}

	addOne := func(slice []int) {
		for i := range slice {
			slice[i]++
		}
	}

	combinedFunc := appendFunc(printSlice, doubleValues, addOne)
	combinedFunc([]int{1, 2, 3, 4})
}
