package main

import "fmt"

//sending to channel
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("processing :", num)
// 		time.Sleep(time.Second)
// 	}
// }



// receiving from the channel
// func sum(result chan int, num1 int, num2 int) {
// 	numRes := num1 + num2
// 	result <- numRes //Sending to a Channel or Send numRes into the result channel.
// }



//goroutine synchronization
func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("processing...")
}




func main() {

	//blocking channel or unbuffered channel
	// msgChannel := make(chan string)
	// msgChannel <- "shubham"
	// msg := <-msgChannel
	// fmt.Println(msg)


	//sending to channel
	// numChan := make(chan int)  //Channel Creation
	// go processNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }


	//receiving from the channel
	// result := make(chan int)   //Channel Creation
	// go sum(result, 4, 5)
	// res := <-result    //Receiving from a Channel
	// fmt.Println(res)

	
	done := make(chan bool)
	go task(done)
	<-done

}

//Main Goroutine

// result := make(chan int)
//         │
//         ▼
// go sum(result,4,5)
//         │
//         ▼
// Main waits
// res := <-result
//              ||
// sum()
// 4 + 5
// ↓
// 9
// ↓
// result <- 9
//              ||
// Main receives
// ↓
// res = 9
// ↓
// Print 9
