package main

import (
	"fmt"
	"sync"
)

func chA(a, b chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("1A")

	<-a
	fmt.Println("A")
	b <- 1
}

func chB(a, b chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("2B")

	<-b
	fmt.Println("B")
	a <- 1
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	cha := make(chan int, 1)
	chb := make(chan int, 1)

	go chA(cha, chb, &wg)
	go chB(cha, chb, &wg)

	cha <- 1
	chb <- 1

	wg.Wait()
}
