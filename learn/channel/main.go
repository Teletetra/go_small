package main

import "fmt"


func processNum(numChan chan int){
	fmt.Println("processing number",<-numChan)
}


func emailSender(emailChan chan string,done chan bool){
	defer func(){done <-true}()
	for email:=range emailChan{
		fmt.Println("sending email to ",email)
	}
}

func main() {

	emailChan:=make(chan string,100)
	done:=make(chan bool)
	for i:=0;i<100;i++{
		emailChan<-fmt.Sprintf("%d@gmail.com",i)
	}
	// numChan:=make(chan int)


	// go processNum()
	// messageChan := make(chan string)

	// messageChan <- "ping" //channel blocking

	// msg:=<-messageChan

	// fmt.Println(msg)
}