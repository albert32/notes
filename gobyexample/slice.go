
package main

import "fmt"
import "strings"

func main(){
    var s = make([]string, 3)
    fmt.Println(s)

    s[0] = "a"
    s[1] = "a"
    s[2] = "a"
  //  s[3] = "a"
    fmt.Println(s)

    t := []string{ "hello", "world"}
    str1 := strings.Join(t, ",")
    fmt.Println(str1)

    twoD := make([][]int,3)
    for i:=0;i<3;i++{
        innerlen := i+1
        twoD[i] = make([]int,innerlen)
        for j:=0;j<innerlen;j++{
            twoD[i][j] = i+j
        }
    }
    fmt.Println(twoD)
}
