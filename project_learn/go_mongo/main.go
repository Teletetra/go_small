package main 

import (
	"fmt"
	// "github.com/julienschmidt/httprouter"
	// "net/http"
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func main(){

	//  we are here using context as kill switch so that it doesnt always open the mongoserver it has a time limit so that it can fetch data so this wrap a base context and add a deadline 

	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()


	client, err := mongo.Connect(
    ctx,
    options.Client().ApplyURI("mongodb+srv://kanishkraval1_db_user:9LLxYbUkB8bZZ6U3@cluster0.9p9dijr.mongodb.net/?appName=Cluster0"),
)
	if err!=nil{
		log.Fatal(err)	
	}

	defer func(){
		if err:=client.Disconnect(ctx);err!=nil{
			log.Println(err)
		}
	}()
	
	if err:= client.Ping(ctx,nil);err!=nil{ 
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	db:=client.Database("testdb")
	users:=db.Collection("users")
	_=users


	r.GET
}

// func getSession() *mgo.Session{
// 	s,err:=mgo.Dial("mongodb://localhost:8089")
// 	if err!=nil{
// 		panic(err)
// 	}
// 	return s
// }