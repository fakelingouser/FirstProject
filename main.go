package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("Test.txt")
	if err != nil {
		fmt.Println("File Test was not created successfully")
	}

	_, err = f.WriteString("This is some text.")
	if err != nil {
		fmt.Println("Failed to write to Test.txt")
	}

	f1, err := os.Open("test1.txt")
	b1 := make([]byte, 5)
	_, err = f1.Read(b1)
	fmt.Println(b1)

	f2, err := os.Open("test2.txt")
	b2 := make([]byte, 5)
	_, err = f2.Read(b2)
	fmt.Println(b2)
}
