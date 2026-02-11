package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go_mongo/models" // Make sure your module name matches go.mod
)

type UserController struct{
	client *mongo.Client 
	collection *mongo.Collection
}


func NewUserController(client *mongo.Client) *UserController{

	collection:=client.Database("testdb").Collection("users")
	return &UserController{client,collection}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request,_ httprouter.Params){

	u:=models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.ID=primitive.NewObjectID()

	ctx,_:= context.WithTimeout(context.Background(),5*time.Second)

	_,err:=uc.collection.InsertOne(ctx,u)
	if err !=nil{
		fmt.Println(err)
	}

	uj,_:= json.Marshal(u)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w,"%s\n",uj)

}

// package controllers

// import ("go.mongodb.org/mongo-driver/bson")


// result,err:=users.InserOne(ctx,bson.M{
// 	"name":"Alice",
// 	"age":25,
// })

// if err!=nil{
// 	log.Fatal(err)
// }

// fmt.Println("Inserted ID:",result.InsertID)


// var user bson.M

// err:=users.FindOne(ctx,bson.M{
// 	"name":"Alice",
// }).Decode(&user)

// // 
// if err!=nil{
// 	log.Fatal(err)
// }

// fmt.Println(user)


// type User struct{
// 	Name string `bson:"name"`
// 	Age int `bson:"age"`
// }

// _,err:=users.InsertOne(ctx,User{
// 	Name:"Bob",
// 	Age:30,
// })