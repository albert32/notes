
package main

import "fmt"

func main(){
    nums := []int{1,2,3,4}
    sum := 0
    for _, num := range nums{
        sum += num
    }

    fmt.Println(sum)

    m := map[string]int {"a":1,"b":2}

    for key, value := range m{
        fmt.Println(key, "->", value)
    }

    for k := range m{
        fmt.Println(k)
    }

}
