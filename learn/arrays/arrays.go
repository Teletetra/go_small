package main

import (
	"fmt"
)

func main() {
	var nums [4]int

	nums[0]=1
	var val[2]bool
	fmt.Println(val)
	num1:=[2][2]int{{3,4},{5,6}}
	fmt.Println((nums[0]))
	fmt.Println(num1)

	// slices
	var slic1 []int
	// slic1[1]=1
	// make func
	var num3 =make([]int,2,5)
	fmt.Println(num3==nil)
	// maximum number of elements can fit 
	slic3:=[]int{}
	slic3=append(slic3, 4)
	fmt.Println(slic3)
	num3=append(num3,1 )
	num3=append(num3,1 )
	num3=append(num3,1 )
	num3=append(num3,1 )
	fmt.Println((num3))
	fmt.Println(cap(num3))
	fmt.Println(slic1)
	copy(num3,slic3)
	fmt.Println(num3,slic3)
}