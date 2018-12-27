
package main

import "fmt"

func plus(a int, b int)int{
    return a+b
}

func plusPlus(a, b, c int)int{
    return a+b+c
}

func myFunc(a int, b int) (int, string){
    return a+b,"hello"
}


func sums(nums ...int) int{
    fmt.Print(nums)
    total := 0
    for _, num := range nums{
        total += num
    }
    fmt.Println(total)
    return total
}


func main(){
    res := plus(1, 2)
    res1 := plusPlus(1, 2, 3)

    fmt.Println(res, res1)
    a,b := myFunc(2,5)
    fmt.Println(a,b)

    sums(1, 2, 3)
    sums(2, 3, 4)

    nums := []int{1, 2, 3, 4, 5}
    sums(nums ...)
}
