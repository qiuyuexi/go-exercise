package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

/**
	尾部插入元素，字底向上调整
 */
func insert(arr []int, length int) {

	for i := length / 2; i > 0; i-- {
		if arr[i] > arr[i*2] {
			swap(&arr[i], &arr[i*2])
		}
		if (i*2+1) <= length && arr[i] > arr[i*2+1] {
			swap(&arr[i], &arr[i*2+1])
		}
	}
	return
}

/**
	弹出头部元素，尾部元素插入头部，数组长度-1 自上向下调整
 */
func pop(arr []int, length int) {
	arr[1], arr[length] = arr[length], arr[1]
	length--
	for i := 1; i < length; i = i * 2 {
		if (i*2) <= length && arr[i] > arr[i*2] {
			swap(&arr[i], &arr[i*2])
		}
		if (i*2+1) <= length && arr[i] > arr[i*2+1] {
			swap(&arr[i], &arr[i*2+1])
		}
	}
}

/**
最小堆
 */
func main() {
	var n int;
	n = 9
	arr := make([]int, n)
	b := []int{0, 9, 3, 7, 6, 5, 1, 10, 2}
	for i := 1; i < len(b); i++ {
		arr[i] = b[i]
		insert(arr, i)
	}
	fmt.Println(arr[1:])
	arr2 := arr
	arr3 :=make([]int,len(arr2))
	for i := len(arr2) - 1; i >= 1; i-- {
		arr3[i] = arr2[1]
		pop(arr2, i)
	}
	fmt.Println(arr3[1:3])
}
