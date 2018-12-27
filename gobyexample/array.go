package main

import "fmt"
import "strconv"

func main(){
    var a [6]int 
    fmt.Println(a)

    b := [5]int{1,2,3,4,5}
    fmt.Println(b)

    var twoD [2][3]int 
    for i:=0;i<2;i++{
        for j:=0;j<3;j++{
            twoD[i][j] = i+j
        }
    }

    fmt.Println(twoD)

    var s [4]string
    for i:=0;i<4;i++{
        s[i] = "hello" + strconv.Itoa(i)
    }
    fmt.Println(s)
}
