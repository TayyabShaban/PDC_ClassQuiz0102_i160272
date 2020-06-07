package main

import (
	"fmt"
)

func main() {
	var no int
	for {
		fmt.Println("enter no 1-4 for country details 0 to quit")
		_, _ = fmt.Scanf("%d", &no)
		if no == 0 {
			break
		} else if no == 1 {
			fmt.Println("Corona cases in Pakistan = 98,943")
		} else if no == 2 {
			fmt.Println("Corona cases in South Korea = 11,776")
		} else if no == 3 {
			fmt.Println("Corona cases in France = 154,000")
		} else if no == 4 {
			fmt.Println("Hello, Corona! ")
		}

	}
}
