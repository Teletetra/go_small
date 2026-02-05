package main

import (
	"fmt"
	"time"
)

func sum() func(int) int{
	sum:=0
	return func(x int)int{
		sum+=x
		return sum
	}
}

type order struct{
	id string
	amount float32
	status string
	createdAt time.Time
}
// structs
// pointers

func main(){
	fmt.Println("Hell")
	x:=sum()
	fmt.Println(x(3))
	fmt.Println(x(4))


	morder:=order{
		id:"1",
		amount: 50.00,
		status:"recieved",
	}
	morder.createdAt=time.Now()
	fmt.Println("this is oru",morder)
}