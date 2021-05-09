package main

import (
	"fmt"
	"sync"
)

/**
问题描述：
使用两个goroutinue 交替打印序列，一个打印数字，另一个打印字母，最终效果如下：
12AB34CD56EF
**/

func printLetter(wait *sync.WaitGroup, letterCh chan bool, numnerCh chan bool) {
	for i := 'A'; i < 'Z'; {
		<-letterCh
		fmt.Printf("%c", i)
		i++
		fmt.Printf("%c", i)
		i++
		numnerCh <- true
	}
	wait.Done()
}

func printNumber(wait *sync.WaitGroup, numberCh chan bool, letterCh chan bool) {
	i := 1
	for {
		select {
		case <-numberCh:
			fmt.Printf("%d", i)
			i++
			fmt.Printf("%d", i)
			i++
			letterCh <- true
		default:
			break
		}
	}
	wait.Done()
}

func main() {
	letter := make(chan bool, 1)
	number := make(chan bool, 1)

	wait := sync.WaitGroup{}
	wait.Add(2)

	number <- true
	go printNumber(&wait, number, letter)
	go printLetter(&wait, letter, number)

	wait.Wait()
}
