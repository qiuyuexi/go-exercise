package main

import (
	"fmt"
	"math/rand"
)

func search(l int, r int, arr []int) int {
	index := l
	for l < r {
		for (l < r) && (arr[r] >= arr[index]) {
			r--
		}
		for (l < r) && (arr[l] <= arr[index]) {
			l++
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	arr[l], arr[index] = arr[index], arr[l]
	return l
}

func sort2(arr []int) {
	if arr[0] > arr[1] {
		arr[0], arr[1] = arr[1], arr[0]
	}
}

func sort3(arr []int) {
	sort2(arr)
	if arr[2] < arr[1] {
		arr[2], arr[1] = arr[1], arr[2]
		sort2(arr)

	}
}

func quickSort(l int, r int, arr []int) {
	if len(arr) == 3 {
		sort3(arr)
		return
	}
	if len(arr) == 2 {
		sort2(arr)
		return
	}
	if l < r {
		index := search(l, r, arr)
		quickSort(l, index-1, arr)
		quickSort(index+1, r, arr)
	}
	return
}

func main() {
	var n int;
	fmt.Scanf("%d,", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Int() % 100
	}
	fmt.Println(arr)
	quickSort(0, n-1, arr)
	fmt.Println(arr)
}
