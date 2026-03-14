package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//go run . --color=red kit "a king kitten have kit"

	data, err := os.ReadFile("../ascii-art/standard.txt")
	if err != nil {
		fmt.Println("Error", err)
	}
	stringSlice := strings.Split(string(data), "\n")
	args := os.Args
	input := args[1]

	fmt.Println(strings.HasPrefix("---color=red", "--"))
	inputSlice := strings.Split(input, "\\n")
	for j := 0; j < len(inputSlice); j++ {
		if inputSlice[j] == "" && j != len(inputSlice)-1 {
			fmt.Println()
		}
		for i := 0; i < 8; i++ {
			for k, r := range inputSlice[j] {
				// fmt.Print(string(r))
				print := ((int(r) - 32) * 9) + i + 1
				fmt.Print(stringSlice[print])
				if k == len(inputSlice[j])-1 {
					fmt.Print("\n")
				}
				// fmt.Print(k)
			}
		}

	}
}
