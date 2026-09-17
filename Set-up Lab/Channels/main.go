package main

import (
	"fmt"
	"time"
)

func main() {
	basicChannels()
	time.Sleep(2 * time.Second)
}

// Channel = pipe, input one end, output the other
// For synchornisation/passing data back and forth
// Use for sharing variables
func basicChannels() {
	//Makes a channel
	c := make(chan int)
	//Make a buffered channel that cna only take 7 ints
	//c := make(chan, int, 7)
	go func() {
		c <- 565656
	}()
	go func() {
		var x int
		var ok bool
		//Read from the c pipe (<- is an arrow)
		//Will stop here until a value is passed into pipe
		x, ok = <-c
		fmt.Println(x)
		//Close the pipe when done (safety hazard, can trip)
		close(c)
		//Initialised to value 0 as pipe is closed
		x, ok = <-c
		if ok == false {
			fmt.Println("Closed, CLOSED...")
			fmt.Println(x)
		}
	}()
}
