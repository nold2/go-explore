package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Panic!")
	}
}

func printStuff() {
	defer wg.Done()
	defer handlePanic()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	wg.Add(1)
	go printStuff()
	go printStuff()
	go printStuff()
	go printStuff()
	go printStuff()
	wg.Wait()
}
