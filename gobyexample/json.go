package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, flag := json.Marshal(1)
	fmt.Println(string(intB), flag)

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	var srcB string
	err := json.Unmarshal(strB, &srcB)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(srcB)

	slcD := []string{"a", "b", "c"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println((slcB))
	fmt.Println(string(slcB))

	mapD := map[string]int{"a": 1, "b": 3, "c": 4}
	fmt.Println(mapD)
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	var dat map[string]interface{}
	error := json.Unmarshal(mapB, &dat)
	fmt.Println(dat, error)

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
