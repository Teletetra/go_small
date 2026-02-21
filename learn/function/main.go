package main

import (
	"fmt"
	"math"
)

// function basic
func add(a int, b int) int {
	return a + b
}
// function with multiple return values 
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return x/y,nil
}
// 3. Named Return Values
func rectangleProp(length,width float64) (area,perimeter float64){
	area=length*width
	perimeter=2*(length+width)
	return
}
// vairiadic function are those which can as many argument as needed using ellipsis(...)

func sum(nums ...int) int{
	total:=0
	for _,num:=range nums{
		total+=num
	}
	return total
}

// anonymous function and closure



// higher order function 
func applyOperation(a,b int ,operation func(int,int) int) int{
	return operation(a,b)
}

//  method function with no recievers
type Circle struct{
	Radius float64
}

func (c Circle) Area() float64{
	return 3.14*c.Radius*c.Radius
}


func (c *Circle) Grow (factor float64){
	c.Radius*=factor
}


// method on non struct type 




type Multiplier int



func (m Multiplier) Multiply(val int) int{

  return int(m)*val



}



// interface in golang 

// how it work 

type Shape interface{

  Area() float64

}



type Circles struct{

  Radius float64

}



type Rectangle struct{

  Width, Height float64

}



func (c Circles) Area() float64{

  return math.Pi*c.Radius*c.Radius

}



func (r Rectangle) Area() float64{

  return r.Width* r.Height

}



func PrintShapeDetails(s Shape){

  fmt.Printf("The area of this shape is:%.2f\n",s.Area())

}



func main(){

  Age:=func(age int)int{

    x:=age+69

    return x

  }

  func (msg string){

    fmt.Println(msg)

  }("executed immeadiately")



  fmt.Println(Age(4))



  counter:=0

  increment :=func()int{

    counter++

    return counter

  }



  fmt.Println(increment())

  fmt.Println(increment())





  addFunc:=func(x,y int) int{return x+y}

  result:=applyOperation(5,7,addFunc)





}

