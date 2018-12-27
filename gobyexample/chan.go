
package main

import "fmt"

var complete chan int = make(chan int)

func f1(){
    for i:=0; i < 10; i++ {
        fmt.Println(i)
    }
    complete <- 1
}

func main(){
    go f1()
    <- complete
}
