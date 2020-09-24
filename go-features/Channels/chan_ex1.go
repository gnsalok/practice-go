package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	//send
	go foo(ch)

	//receive
	bar(ch)

}

//send
func foo(ch chan<- int) {
	ch <- 33

}

// receive
func bar(ch <-chan int) {
	fmt.Println(<-ch)

}
