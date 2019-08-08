package main

import ("fmt"
"sort"
)


func main(){
	n := []int {4,6,3,8,2,9}

	fmt.Println("normal:", n)

	sort.Ints(n)

	fmt.Println("ascendente:", n)
}
