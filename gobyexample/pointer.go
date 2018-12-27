package main

import "fmt"

func zeroval(val int){
    val = 0
}

func ptrval(val *int){
    *val = 0
}

func main(){
    i := 1
    fmt.Println(i)
    zeroval(i)
    fmt.Println(i)

    ptrval(&i)
    fmt.Println(i)
    fmt.Println(&i)
}
