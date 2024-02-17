package main

import "fmt"

func longestPalindrome(s string) string {

	max_arr := make([]int, len(s))
	max_str := ""
	for i := range max_arr {
		max_arr[i] = 1
	}

	for i, _ := range s {
		left, right := i-1, i+1
		max_pal := string(s[i])
		for left >= 0 && right < len(s) && s[left] == s[right] {
			max_pal = string(s[left]) + max_pal + string(s[right])
			left--
			right++
		}
		max_arr[i] = len(max_pal)
		if len(max_pal) > len(max_str) {
			max_str = max_pal
		}

		if i+1 < len(s) && s[i] == s[i+1] {
			left, right := i-1, i+2
			max_pal := string(s[i : i+2])
			for left >= 0 && right < len(s) && s[left] == s[right] {
				max_pal = string(s[left]) + max_pal + string(s[right])
				left--
				right++
			}
			max_arr[i] = max(len(max_pal), max_arr[i])
			if len(max_pal) > len(max_str) {
				max_str = max_pal
			}
		}
	}
	// fmt.Println(max_arr)
	return max_str
}

func main() {

	s := "aaaaaaa"
	fmt.Println(longestPalindrome(s))
}
