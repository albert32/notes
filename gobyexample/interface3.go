
package main

import "fmt"

type I interface{
    Get() int 
    Set(int)
}

type S struct{
    value int 
}

func (s S) Get() int{
    return s.value
}

func (s *S) Set(a int){
    s.value = a
}

func main(){
    s := S{}

    var i I

    i = &s

    i.Set(5)
    fmt.Println(i.Get())

    switch v := i.(type){
        case *S:
            fmt.Println(v)
            fmt.Println("ssssss")
        default:
            fmt.Println("ttttt")
    }

}
