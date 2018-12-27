
package main

import (
    "fmt"
    "time"
)

func cal(a int, b int){
    c := a + b
    fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main(){
    for i:=0; i<10; i++{
        go cal(i, i+1)
    }

    time.Sleep(time.Second * 2)
}
