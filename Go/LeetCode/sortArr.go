package main

import (
	"fmt"
	"sort"
)

func returnMedian(arr []int) float32 {
	middle := len(arr)/2 - 1

	if len(arr)%2 == 0 {
		fmt.Println("Even")
		return (float32(arr[middle]) + float32(arr[middle+1])) / 2
	}
	return float32(arr[middle])
}

func main() {

	arr1 := []int{1, 2, 6, 7}
	arr2 := []int{5, 6, 7, 2}

	m := len(arr1)
	n := len(arr2)

	var arr3 []int

	sort.Ints(arr1)
	sort.Ints(arr2)

	i := 0
	j := 0

	for i < m && j < n {
		if arr1[i] < arr2[j] {
			arr3 = append(arr3, arr1[i])
			i++
		} else {
			arr3 = append(arr3, arr2[j])
			j++
		}
	}

	for i < m {
		arr3 = append(arr3, arr1[i])
		i++
	}
	for j < n {
		arr3 = append(arr3, arr2[j])
		j++
	}

	fmt.Println(m + n)
	fmt.Println(arr3)
	fmt.Println("The meadian is ", returnMedian(arr3))
}
