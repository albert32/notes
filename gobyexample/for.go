
package main

import "fmt"

func main(){
    for {
        fmt.Println("loop.")
        break
    }

    i := 1
    for i < 3 {
        fmt.Println(i)
        i += 1
    }

    for j := 1; j < 5; j++ {
        fmt.Println(j)
    }

}
