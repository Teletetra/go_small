package main

import "fmt"

type payment struct {
	gateway razorpay

}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw :=razorpay{}
	// stripePaymentGw:=stripe{}
	// razorpayPaymentGw.pay(amount)
	p.gateway.pay(amount)
}


type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay",amount)
}

type stripe struct{}

func (s stripe) pay(amount float32){
	fmt.Println("making payment using stripe",amount)
}

type fakepayment struct {}


func (f fakepayment) pay (amount float32){
	fmt.Println("making payment using fake gateway for testing purpose")
}


// interface introuduce

// convention name-er

type paymenter interface{
	pay(amount float32)
}

func main() { 
	// stripePaymentGw:=stripe{}
	razorpayPaymentGw:=razorpay{}

	payment1:=payment{
		gateway: razorpayPaymentGw,
	}
	payment1.makePayment(100)
}