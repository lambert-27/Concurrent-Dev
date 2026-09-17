package main

import (
	"fmt"
	"time"
)

// TIP To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	//TIP Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s \n", s)
	fmt.Println("Hello ", s)
	race()
	fmt.Println("Next")
	//Start with capital letter = public (belongs to FMT)
	//Start with lowercase = private
	time.Sleep(2000 * time.Millisecond)
	threads(5)
	time.Sleep(2000 * time.Millisecond)

	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}

}

func threads(num int) {
	for i := 0; i < num; i++ {
		// Anonymous function - has no name
		//go func(num int) {
		//	fmt.Printf("this is thread: %d \n", num)
		//}(i)
		//Can be done this way either, go keyword sends straight to func
		//Sends printy to do its own exe while main continues
		go printy(i)
		fmt.Println("thread launched", i)
	}
}

func printy(num int) {
	fmt.Printf("this is thread: %d \n", num)
}

// Race condition not guaranteed
func race() {
	i := 0
	//Map/array of strings
	//_ = index, s=value, [] = unsized array
	for _, s := range []string{"a", "b", "c", "d", "e"} {
		i = i + 1
		//Anonymous function
		go func(num int, letter string) {
			fmt.Printf("Thread: %d letter: %s \n", num, letter)
		}(i, s)
		fmt.Println("launched", i)
	}
}

// w/ race condition guarunteed
func raceC() {
	i := 0
	//Map/array of strings
	//_ = index, s=value, [] = unsized array
	for _, s := range []string{"a", "b", "c", "d", "e"} {
		i = i + 1
		//Anonymous function
		go func(num int) {
			fmt.Printf("Thread: %d letter: %s \n", num, s)
		}(i)
		fmt.Println("launched", i)
	}
}
