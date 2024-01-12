package main

import "fmt"

/*
    1
	2 2
	3 3 3
	4 4 4 4
	5 5 5 5 5
*/

func main(){
	for i:=1;i<=5;i++{
		for j:=1; j<=i; j++{
			fmt.Print(i)
		}
		fmt.Println()
	}
}