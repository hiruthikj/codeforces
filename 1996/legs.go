package main

import "fmt"

func main() {
	var t, n int
	fmt.Scanf("%d\n", &t)


	for i:=0; i<t; i++ {
		fmt.Scanf("%d\n", &n)	
		minLegs := n/4 + (n%4)/2
		fmt.Println(minLegs)
	}


}
