package main

import (
	"ascii-art-color/operations"
	"fmt"
)

func main() {

	ok, input := operations.ValidateInput()

	if !ok {
		return
	}

	//Check if there is no input at all eg when len(args) != 2
	if len(input.Str) == 0 {
		return
	}

	//check if there is only "" in string
	if input.Str[0] == "" && len(input.Str) == 1 {
		return
	}

	// //check if 
	if !ok {
		return
	}

	//Check if therinput only contain new line ie; "\n"
	if input.Str[0] == "\n" && len(input.Str) == 1 {
		fmt.Println()
		return
	}

	data := operations.ReadTextFile("standard.txt")

	//Check if there is no input at all eg when len(args) != 2 or when there is an
	//error in ReadFile
	if len(input.Str) == 0 || len(data) == 0 {
		return
	}
	// inputSlice := strings.Split(input, "\\n")
	result := operations.AsciiArt(input, data)

	fmt.Print(result)
}
