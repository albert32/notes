
package main

import "fmt"

func main(){
    s := "hello 你好"
    fmt.Println("len is",len(s))
    fmt.Println("rune:", len([]rune(s)))
}
