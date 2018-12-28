package main

import (
	"fmt"
	"os"
)

func main() {
	file := createFile("data.txt")
	defer closeFile(file)
	writeFile(file, "hello world")
}

func createFile(filename string) *os.File {
	fmt.Println("start create file")
	file, error := os.Create(filename)
	if error != nil {
		panic(error)
	}
	return file
}

func writeFile(file *os.File, str string) {
	fmt.Printf("write %v to file \n", str)
	fmt.Fprintln(file, str)
}

func closeFile(file *os.File) {
	fmt.Println("close file")
	file.Close()
}
