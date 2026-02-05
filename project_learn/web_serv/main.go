package main

import (
	"fmt"  // printing text (to browser & console)
	"log"
	"net/http" //HTTP server, routing, requests, responses
) 

// Efficiency: An http.Request is a large struct containing many fields (headers, URL, cookies, etc.). Passing it by value would force Go to copy the entire struct into memory for every handler call, which is expensive. A pointer is just a small memory address (typically 8 bytes), making it much faster to pass around.


// Shared State & Mutability: Certain methods like r.ParseForm() or r.Context() modify the state of the request. By using a pointer, these updates are made to the original request rather than a local copy, ensuring that subsequent middleware or handlers see the updated data (like parsed form values). 
// Why is ResponseWriter NOT a pointer?
// You might notice http.ResponseWriter is passed as a value (no *). This is because ResponseWriter is an interface, not a struct. Under the hood, the data satisfying that interface is already a pointer, so there is no need for an extra 
func formHandler(w http.ResponseWriter,r *http.Request){
	//  ParseForm basically read data from the http request work url quert param ,post from data. it retrun nil value if everything go along the way.
	// and then short variable declaration inside if , it stores the return value of r.ParseForm() in err and the scope of this variable is inside of this if block.  
	if err:=r.ParseForm();err!=nil{
		fmt.Fprintf(w,"ParsedForm() err:%v",err)
		return
	}
	// if r.Method=="POST"{
	// 	fmt.Fprintf(w,"post request succesful\n")
	// 	name:=r.FormValue("name")
	// 	address:=r.FormValue("address")
	// 	fmt.Fprintf(w,"Name=%s\n",name)
	// 	fmt.Fprintf(w,"address=%s\n",address)
	// }
	if r.Method=="POST"{
		fmt.Fprintf(w,"post request successfull\n")
		name:=r.FormValue("name")
		address:=r.FormValue("address")
		mobile:=r.FormValue("mobile")
		fmt.Fprintf(w,"Name=%s\n",name)
		fmt.Fprintf(w,"Address=%s\n",address)
		fmt.Fprintf(w,"Mobile=%v\n",mobile)
	}
}
// t uses a format string (containing "verbs" like %s for strings or %d for integers) to determine exactly how the output should look.
func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path!="/hello"{
		http.Error(w,"404 Not Found",http.StatusNotFound)
		return
	}
	// if r.Method!="GET"{
	// 	http.Error(w,"method is not supported",http.StatusNotFound)
	// 	return
	// }
	if r.Method!=http.MethodGet{
		http.Error(w,"method not allowed",http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(w,"hello")	
}

func main(){
	// fileServer:=http.FileServer(http.Dir("./static"))
	
	fileServer:=http.FileServer((http.Dir("./static")))
	// http.Handle("/",fileServer)
	// http.HandleFunc("/form",formHandler)

	// http.HandleFunc("/hello",helloHandler)

	// fmt.Println("starting server at port 8080")
	// if err:=http.ListenAndServe(":8080",nil);err !=nil{
	// 	log.Fatal(err)
	// }
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)
	fmt.Println("starting server at port 8082")
	if err:=http.ListenAndServe(":8082",nil);err!=nil{
		log.Fatal(err)
	}
	age:=map[string]int{
		"kanishk":23,
	}
}