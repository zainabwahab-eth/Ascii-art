package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error", err)
	}

	stringSlice := strings.Split(string(data), "\n")
	// fmt.Println(stringSlice[9])
	// fmt.Println(stringSlice[10])
	// fmt.Println(stringSlice[11])
	// fmt.Println(stringSlice[12])
	// fmt.Println(stringSlice[13])
	// fmt.Println(stringSlice[14])
	// fmt.Println(stringSlice[15])
	// fmt.Println(stringSlice[16])
	// fmt.Println(stringSlice[17])

	// (ascii - 32)*9
	// (65 - 32)*9

	args := os.Args
	input := args[1]

	inputSlice := strings.Split(input, "\\n")
	// fmt.Println(len(inputSlice))

	// fmt.Println(inputSlice[1] == "")
	// if strings.Contains(input, "\n") {
	// }

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
		// if j != len(inputSlice) - 1{
		// 	fmt.Print("\n")
		// }
		// fmt.Print("\n")
	}
	//////////////////////////////////////VERSION 1

	// for i := 0; i <= 8; i++ {
	// 	for k, r := range input {
	// 		print := ((int(r) - 32) * 9) + i + 1
	// 		fmt.Print(stringSlice[print])

	// 		if k == len(input)-1 {
	// 			fmt.Print("\n")
	// 		}
	// 	}
	// }
}
