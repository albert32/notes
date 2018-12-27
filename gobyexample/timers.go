
package main

import "fmt"
import "time"

func main(){
    timer1 := time.NewTimer(2 * time.Second)
    <- timer1.C
    fmt.Println("timer1 expired")

    timer2 := time.NewTimer(time.Second)
    go func(){
        <-timer2.C
        fmt.Println("timer2 expired")
    }()

    stop := timer2.Stop()
    if stop{
        fmt.Println("stop timer2")
    }
}
