package main

import (
	"fmt"
)

func main(){
	fmt.Println("howdi partner") // just printing
	var p *int
	var q *int
	j := 7
	p = &j
	q = p+1
	fmt.Println(q) //pointer operations
}