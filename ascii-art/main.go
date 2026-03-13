package main

import (
	"ascii-art/operations"
	"fmt"
)

func main() {
	input := operations.GetInput()

	//Check if there is no input at all eg when len(args) != 2
	if len(input) == 0 {
		return
	}

	//check if there is only "" in string
	if input[0] == "" && len(input) == 1  {
		return
	}
	
	// //check if input only contain new line ie; "\n"
	if input[0] == "\n" && len(input) == 1 {
		fmt.Println()
		return
	}
	
	data := operations.ReadTextFile("standard.txt")
	
	//Check if there is no input at all eg when len(args) != 2 or when there is an 
	//error in ReadFile
	if len(input) == 0 || len(data) == 0 {
		return
	}
	// inputSlice := strings.Split(input, "\\n")
	result := operations.AsciiArt(input, data)

	fmt.Print(result)
}
