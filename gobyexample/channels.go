package main

import "fmt"
import "time"

func main(){
    messages := make(chan string, 3)

    go func(){
        messages <- "hello world"
        messages <- "hello world"
        messages <- "hello world"
        close(messages)
    }()

  //  close(messages)
    for msg := range messages{
        fmt.Println(msg)
       //   if len(messages) <= 0{
          //  break
        //}
    }
    time.Sleep(time.Second * 5)

	time.Sleep(time.Second * 5)
}

