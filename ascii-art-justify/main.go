package main

import (
	"ascii-art-justify/operations"
	"fmt"
	"strings"
)

func printString(align string, resultSlice [][]string, totalWidth int, output string) {
	terminalWidth, err := operations.TerminalWidth()
	if err != nil {
		return
	}
	padding := terminalWidth - totalWidth

	switch align {
	case "right":
		operations.PrintRight(resultSlice, padding)

	case "centre":
		operations.PrintCentre(resultSlice, padding)

	case "justify":
		operations.PrintJustify(resultSlice, padding)

	default:
		result := operations.PrintDefault(resultSlice)
		// fmt.Println("res", result)

		if output == "" {
			fmt.Print(result)
		} else {
			// printRes := result
			operations.WriteFile(result, output)
		}
	}

}

func main() {
	ok, input := operations.ValidateInput()

	if !ok {
		return
	}

	strSlice := strings.Split(input.Str, "\\n")

	//Check if there is no input at all eg when len(args) != 2
	if len(strSlice) == 0 {
		return
	}

	//check if there is only "" in string
	if strSlice[0] == "" && len(strSlice) == 1 {
		return
	}

	// //check if input only contain new line ie; "\n"
	if strSlice[0] == "\n" && len(strSlice) == 1 {
		fmt.Println()
		return
	}

	banner := "standard.txt"
	if input.Banner != "" {
		banner = input.Banner + ".txt"
	}
	err, data := operations.ReadTextFile(banner)

	if err != nil {
		fmt.Println("error", err)
		return
	}
	//Check if there is no input at all eg when len(args) != 2 or when there is an
	//error in ReadFile
	if len(input.Str) == 0 || len(data) == 0 {
		return
	}

	//loop through string gotten from spliting by "\n"
	for _, s := range strSlice {
		resultSlice, totalWidth := operations.PerformAscii(s, input, data)
		printString(input.Alignment, resultSlice, totalWidth, input.Output)
	}

}
