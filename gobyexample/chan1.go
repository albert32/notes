
package main

import "fmt"
import "time"

func f(done chan bool){
    fmt.Println("start...")
    for i:=0;i<10;i++{
        fmt.Println(i)
    }
    time.Sleep(time.Second * 5)
    fmt.Println("end...")
    done <- true
}

func main(){
    done := make(chan bool, 1)

    go f(done)

    <- done
}
