package main

import ("fmt"
"sort"
)


func main(){
	n := []int {4,6,3,8,2,9}

	fmt.Println("normal:", n)

	sort.Ints(n)
        
	var frontend []int = [1:4]
	
	fmt.Println("ascendente:", frontend)
}
