package main 

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"gopkg.in/mgo.v2"
)

func main(){
	r:=httprouter.New()
	uc:=controllers.NewUserController(getSession())
	r.GET("/user/:id",uc.GetUser)
	r.POST()
	r.DELETE()
}

func getSession() *mgo.Session{
	s,err:=mgo.Dial("mongodb://localhost:8089")
	if err!=nil{
		panic(err)
	}
	return s
}