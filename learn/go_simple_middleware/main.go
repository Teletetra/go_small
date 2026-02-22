package main

import (
	"fmt"
	"time"
)

type HandlerFunc func(request string)

func TimerMiddleware(next HandlerFunc) HandlerFunc{
	return func(request string){
		start:=time.Now()
		next(request)
		fmt.Printf("Request '%s' took %v n",request,time.Since(start))
	}
} 

func main (){
	processUser:=func(req string){
		fmt.Println("Processing user ")
		time.Sleep(1*time.Second)
	}
	timedProcessUser:=TimerMiddleware(processUser)
	
	timedProcessUser("/api/users")
}