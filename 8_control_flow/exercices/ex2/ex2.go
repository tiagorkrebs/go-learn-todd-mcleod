package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j <= 2; j++ {
			fmt.Printf("\t%U %c\n", i, i)
		}
	}
}
