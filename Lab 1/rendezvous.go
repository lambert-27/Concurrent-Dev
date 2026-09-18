package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// --------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Mark Lambert C00192497
// c00192497@setu.ie
// 18/09/2026
// -------------------------------------------
// Global variables shared between functions --A BAD IDEA
var aArrived = make(chan struct{})
var bArrived = make(chan struct{})

func WorkWithRendezvous(wg *sync.WaitGroup, Num int) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))

	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	//Rendezvous here
	// a sends signal
	aArrived <- struct{}{}

	// b waits
	<-bArrived
	fmt.Println("PartB", Num)

	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	threadCount := 5

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N)
	}

	// All a's wait for b
	for range threadCount {
		<-aArrived
	}

	// Signal all bArrived for a
	for range threadCount {
		bArrived <- struct{}{}
	}

	wg.Wait() //wait here until everyone (10 go routines) is done
}
