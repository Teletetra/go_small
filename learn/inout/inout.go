package main
import "fmt"

func add(a,b int)int{
	return a+b

}

func mul()(string,string){
	return "mul","mul "
}

func process(fn func(a int)int){
	 fn(1)
}

func process2() func(a int) int{
	return func(a int)int{
		return 4
	}
}

// variadic function -are those whicch ccan multiple parameters in the function -they are like spread operator 

func sum(nums ...int)int{
	total:=0
	for _,num:=range nums{
		total=total+num
	}
	return total
}
func sum1(nums ...interface{})int{
	total:=0
	for _,num:=range nums{
		fmt.Println(num)
	}
	return total
}
func main() {
	var name string
	result:=sum(3,4,5,5)
	sli:=[]int{1,4,5,5}
	result2:=sum(sli...)
	fmt.Scanln(&name)
	fmt.Println(name)
	x:=add(3,5)
	fmt.Println(x)
}
