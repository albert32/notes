package main

import "fmt"

func Swap(a, b string) (string, string){
    return b, a
}

func main(){
    c, d := Swap("hello", "world")
    fmt.Println(c, d)
}
