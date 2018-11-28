package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1,4,6,8,9,5,3}
	sort.Ints(a)
	for k,_ := range a{
		fmt.Println(a,k)
	}
}
