package operations

import (
	"os"
	"strings"
)

func GetInput() []string{
	//Get args from os.args
	args := os.Args

	//check that user don't input more than one input
	if len(args) != 2 {
		return []string{}
	}

	//get input from args
	input := args[1]

	//check if there is only new line input
	if input == "\\n" {
		return []string{"\n"}
	}

	//Split input by new line (\n)
	inputSlice := strings.Split(input, "\\n")

	return inputSlice
}

func ReadTextFile(file string) []string {
	//Use os.readfile to read our banner file
	data, err := os.ReadFile(file)

	//Return an empty slice if there is an error
	if err != nil {
		return []string{}
	}

	//Split data by new line
	dataSlice := strings.Split(string(data), "\n")

	return dataSlice
}

func AsciiArt(inputSlice []string, dataSlice []string) string {

	var builder strings.Builder

	//Loop through input array
	for _, input := range inputSlice {

		if input == "" {
			builder.WriteString("\n")
			continue
		}
		
		//Loop 8 times to print all rowa of input
		for row := 1; row <= 8; row++ {

			//Now for each loop above print that row for each chs of input eg
			//row = 1 print row 1 of all chs of input 
			for _, ch := range input {

				//Make sure ch is printable
				if ch >= 32 && ch <= 126 {

					//Calculate formular to get each rows
					print := ((int(ch) - 32) * 9) + row

					//Store string in buffer
					builder.WriteString(dataSlice[print])
				}

			}
			builder.WriteString("\n")
		}
	}

	return builder.String()
	
}
