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
}
