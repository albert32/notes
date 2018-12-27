
package main

import "fmt"

func f(n int) int{
    if n == 0 {
        return 1
    }
    return n * f(n-1)
}

func main(){
    num := f(20)
    fmt.Println(num)
    fmt.Println(f(2))
}
