
package main

import (
    "fmt"
)

func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"stanley", "david", "oscar"}
    vals := make([]interface{}, len(names))
    for i, v := range names {
        vals[i] = v
    }
    PrintAll(vals)

    nums := []int{1, 2, 3, 4}
    nums2 := make([]interface{}, len(nums))
    for i, v := range nums{
        nums2[i] = v
    }
    PrintAll(nums2)
}
