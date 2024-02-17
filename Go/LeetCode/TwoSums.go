package main

import "fmt"

func main() {
	var array [4]int
	array[0] = 3
	array[3] = 4

	fmt.Println("Array : ", array)

	var int_map map[int]int

	for i := 0; i < len(array); i++ {
		int_map[i] = array[i]
	}

	fmt.Print(int_map)

}
