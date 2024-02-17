package main

import "fmt"

func longLeft(s []rune, i int, max_arr []int) (int, error) {

	r := 1
	if i+1 >= len(s) {
		return r, nil
	}
	if s[i] == s[i+1] {
		r = 2
	}
	if i+2 < len(s) && s[i] == s[i+2] {
		r = 3
	}
	if i+max_arr[i+1]+1 < len(s) && s[i] == s[i+max_arr[i+1]+1] {
		r = max(r, max_arr[i+1]+2)
	}
	fmt.Println(i, max_arr[i+1], i+max_arr[i+1], r)
	return r, nil
}
func longestPalindrome(s string) string {
	// c substing of s is palindrome if
	// len(c)==1
	// len(c)==2 c[0]==c[1]
	// c[0] == c[-1] and c[1:-1] is a palindrome
	max_arr := make([]int, len(s))
	r := []rune(s)
	max_index := 0
	for i := range max_arr {
		max_arr[i] = 1
	}

	for index, _ := range r {
		inv_index := len(r) - index - 1

		fmt.Println(inv_index, max_arr)
		pal, _ := longLeft(r, inv_index, max_arr)
		if pal > max_arr[max_index] {
			max_index = inv_index
		}
		max_arr[inv_index] = pal
	}
	fmt.Println(max_index, max_arr)
	return s[max_index : max_index+max_arr[max_index]]
}

func main() {

	s := "aaaaaa"
	fmt.Println(longestPalindrome(s))
}
