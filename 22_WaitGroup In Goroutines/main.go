// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func printHello(wg *sync.WaitGroup) {

// 	defer wg.Done() //Done() means This goroutine has finished.
// 					//Done() is equivalent to Add(-1)

// 	fmt.Println("Hello")
// }

// func main() {

// 	var wg sync.WaitGroup  //Create a WaitGroup.Think of it like a counter=0.

// 	wg.Add(1) //One goroutine is going to start.Counter=1

// 	go printHello(&wg) //Start the goroutine.

// 	wg.Wait()  //It says  Wait...  until counter=0 Only then continue.

// 	fmt.Println("Main Finished")
// }

package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task", id)
}

func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
	fmt.Println("All Tasks Finished")

}

//Easy Way to Remember
// WaitGroup
// ↓
// Counter
// Add()
// ↓
// Counter = +1
// Done()
// ↓
// Counter = -1
// Wait()
// ↓
// Pause until Counter = 0
// ↓
//Continue Program
