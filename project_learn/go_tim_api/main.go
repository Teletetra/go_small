package main

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In the search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "search of Lost ", Author: "meena Proust", Quantity: 2},
	{ID: "3", Title: "Lost Time", Author: "albert Proust", Quantity: 2},
}


func main(){
	router
}