package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float64())

	s2 := rand.NewSource(12)
	r2 := rand.New(s2)

	fmt.Println(r2.Intn(22))
	fmt.Println(r2.Float64())
}
