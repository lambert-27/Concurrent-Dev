package main

import (
	"fmt"
	"sync"
)

//A closed pipe never blocks - Joe

//Create lock, and lock it (m)
//If defer function runs first, it will stop because it tries to lock an already locked mutex

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select
func main() {
	var m sync.Mutex
	var a int //what is value of a
	done := make(chan struct{})
	//Mutex locked (no variable in mutex locked, mutex is empty)
	m.Lock()

	//	how is defer used? (and why?)
	go func() {
		//Blocks, as it is already locked, waits for it to be unlocked
		//Think of bathroom door, we lock it as soon as we get in, we wait if it is alreayd locked
		m.Lock()
		//Says doo this @ the end of the scope (last thing that happens before the function exits)
		//Makes sure lock and unlock occur together in the code, making it easier to find
		defer m.Unlock()
		fmt.Println(a)
		//m.Unlock()
		//Closed channel returns, channel dies (returns default value and ok bit = false)
		close(done)
	}()

	//We want this function to run before the above one
	go func() {
		//Not locked, assigns value
		a = 1
		m.Unlock()
	}()

	//What does this do? Why have a channel?
	<-done
	//	Channel stops it from exiting, blocks forever, unless we call close
}
