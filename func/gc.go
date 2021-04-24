package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := getRandom()
	fmt.Println(*num)
}

func getRandom() *int {
	tmp := rand.Intn(100)
	return &tmp
	}
