package main

import "fmt"

// type OrderStatus int
// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

type OrderStatus string
const (
	Received  OrderStatus = "received" //Received OrderStatus = ""
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to:-", status)
}

func main() {
	//changeOrderStatus(Prepared)
	changeOrderStatus(Prepared)
	fmt.Println()

}
