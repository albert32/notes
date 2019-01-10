package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	args1 := args[1:]
	fmt.Println(args)
	fmt.Println(args1)
}
