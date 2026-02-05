package main

import (
	"fmt"
	"time"
)

func changeNum(num *int) {
	*num = 6
	fmt.Println("In changenum", *num)

}
// interfaces
type payment struct{

}

func (p payment) makePayment(amount float32){
	// rajorpayPaymentG:=rajorpay{}
	// rajorpayPaymentG.pay(amount)
	stripePayment:=stripe{}
	stripePayment.pay(amount)
}

type rajorpay struct{}
type stripe struct{}


func (s stripe) pay(amount float32){
	fmt.Println("making payment using stripe",amount)
}
func (r rajorpay) pay(amount float32){
	fmt.Println("making payment using rajorpay",amount)
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}
// making constructor

func newOrder(id string,amount float32, status string) *order{
	// initial setup if library want to use
	myorder:=order{
		id:id,
		amount: amount,
		status: status,
	}
	return &myorder
}


// reciever type
func (o *order) changeStatus(status string) {
	o.status = status
}

// struct embedding
type customer struct{
	name string
}
func main() {
	// num := 1
	newpayment:=payment{}
	newpayment.makePayment(1090)

	// fmt.Println("after changenum", num)
	// changeNum(&num)
	// fmt.Println("after changenum", num)
	// newCustomer:=customer{"john"}
	// myorder := order{
	// 	id:     "1",
	// 	amount: 4.44,
	// 	status: "recieved",
	// 	customer: newCustomer,
	// }
	// myorder.changeStatus("not yet")
	// fmt.Println(myorder)
}
