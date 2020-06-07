package main

import (
	"Quiz1pkg"
	"fmt"
)

func main() {
	var no int
	for {
		fmt.Println("enter no 1-4 for country details 0 to quit")
		_, _ = fmt.Scanf("%d", &no)
		var ans string = Quiz1pkg.PrintCoronaDetails(no)
		if ans == "break" {
			break
		} else {
			fmt.Print(ans)
		}

	}

}
