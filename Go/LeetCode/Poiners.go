package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	rand_num := 3
	ptr := &rand_num

	fmt.Println("value : ", rand_num)
	fmt.Println("address : ", &rand_num)
	fmt.Println("pointer value : ", ptr)
	fmt.Println("pointer address : ", &ptr)
	fmt.Println("value stroed in Ptr : ", *ptr)
	fmt.Println("Doing math ", 7%2, 7/2)

	l1 := ListNode{}
	stptr := &l1

	for i := 0; i < 10; i++ {
		stptr.Val = i * 2
		stptr.Next = &ListNode{}
		stptr = stptr.Next
	}

	stptr = &l1

	a, b := 1, 2

	fmt.Print(a, b)

	// charMap := make(map[string]int)
	var s string = "abcdef"
	fmt.Println("Printint String : ", string(s[2]))
	fmt.Println(max(0, 1))

}
