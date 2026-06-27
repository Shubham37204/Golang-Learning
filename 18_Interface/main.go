package main

import "fmt"

type paymenter interface {
	Pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.Pay(amount)
}

type razorpay struct{}

func (r razorpay) Pay(amount float32) {
	fmt.Println(amount)
}

type stripe struct{}

func (s stripe) Pay(amount float32) {
	fmt.Println(amount)
}

type paypal struct{}
func (s paypal) Pay(amount float32) {
	fmt.Println(amount)
}
func (s paypal) refund(amount float32, account string) {
	fmt.Println(amount, account)
}


func main() {
	//stripePaymentGw := stripe{}
	//razorPaymentGw := razorpay{}
	paypaymentGw := paypal{}
	newPay := payment{
		//gateway: razorPaymentGw,
		//gateway: stripePaymentGw,
		gateway: paypaymentGw,
	}
	newPay.makePayment(100)
	fmt.Println()

}
