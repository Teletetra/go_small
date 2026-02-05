package main

import (
	"fmt"
	"time"
)

// import "rsc.io/quote"

func main() {

	fmt.Println(1 + 1)
	fmt.Println("hey gaurav u are recovering,things just take time okay ")
	var name string = "kanishk and gaurav"
	fmt.Println(name)
	// fmt.Print(quote.Go())
	age := 22
	fmt.Println("age of kanishk", age, "still young but passionate")

	var arr [3]int = [3]int{1, 2, 3}

	fmt.Println(arr)

	slicek := []int{1, 3, 4}
	fmt.Println(slicek)
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println(m)
	expr := (1 + 3) * 4 / 3
	fmt.Println(expr, "value of expr")
	if n := 6; n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
	for i, v := range m {
		fmt.Println(i, v)
	}
	// while loop with for
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	// for {} infinite loop
	for i := range 3 {
		if i > 0 && i == 2 {
			fmt.Println(i)
		}
	}
	// day := "monday"
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("start of hell after wastage")
	case time.Friday:
		fmt.Println("start of hell after wastage")
	}
	arr1 := []int{1, 3, -4, 5, -9}
	for _, k := range arr1 {
		switch {
		case k > 0:
			fmt.Println("the number is positve")
		case k < 0:
			fmt.Println("the number is negative")
		}
	}
whoAMI:=func(i interface{}){
	switch t:=i.(type) {
	case int:
		fmt.Println("dfsd",t)
	}
}
whoAMI("golang")
}
