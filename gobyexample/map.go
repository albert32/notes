
package main

import "fmt"

func main(){
    var m = make(map[string]int)

    m["a"] = 222
    fmt.Println(m)
    m["b"] = 333

    fmt.Println(len(m))

    m1 := map[string]int{"a":1,"b":2}
    fmt.Println(m1)
    
    delete(m1, "a")
    fmt.Println(m1)
    res,exist := m1["a"]
    fmt.Println(res,exist)

    res,exist = m1["b"]
    fmt.Println(res,exist)
}
