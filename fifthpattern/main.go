package main

import "fmt"

/*
 *****
 ****
 ***
 **
 *
 */

func main(){
	for i:=5;i>=1;i--{
		for j:=1;j<=i;j++{
			fmt.Print(" * ")
		}
		fmt.Println()
	}
}
 //second way of approach 5--rows 
// for i:=1;i<=5;i++{
// 	for j:=0;j<=5-(i+1);j++{
// 		fmt.Print(" * ")
// 	}
// 	fmt.Println()
// }

