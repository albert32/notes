
package main

import "fmt"

type person struct{
    name string
    age int
}

func main(){
    fmt.Println(person{"albert", 20})   
    fmt.Println(person{name:"albert", age:20})   
    fmt.Println(person{name:"albert"})   

    s := person{"albert",20}
    t := &s

    fmt.Println(t.name, t.age)
}
