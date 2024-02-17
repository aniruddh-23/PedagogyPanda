package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Userinput")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating : ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks, ", input)

}
