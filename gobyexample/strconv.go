package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseFloat("2.122", 64))
	fmt.Println(strconv.ParseInt("123", 0, 64))
	fmt.Println(strconv.Atoi("222"))
}
