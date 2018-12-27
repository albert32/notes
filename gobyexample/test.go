package main

import (
    "fmt"
)

//http://www.cnblogs.com/osfipin/
func main() {
    var s = "gohahdh"
    //var r = []rune(s)
    for i:=0;i<len(s);i++{
        fmt.Printf('%c',s[i])
        }
    fmt.Println(len(s))
   // fmt.Println(len(r))

    //fmt.Printf("%c ", s[3])
    //fmt.Printf("%c\n", r[3])
}

