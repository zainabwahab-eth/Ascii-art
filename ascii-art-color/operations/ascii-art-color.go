package operations

import (
	"os"
	"slices"
	"strings"
)

type ColorStruct struct {
	Start []int
	End   []int
}

func CheckColor(input string, subString string) (bool, ColorStruct) {
	count := strings.Count(input, subString)
	start := []int{}
	end := []int{}

	if count > len(input) {
		return false, ColorStruct{}
	}

	in := strings.LastIndex(input, subString)

	for in >= 0 {
		start = append(start, in)
		end = append(end, in+len(subString)-1)
		input = input[:in+len(subString)-1]
		in = strings.LastIndex(input, subString)
	}
	return true, ColorStruct{
		Start: start,
		End:   end,
	}

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

func AsciiArt(inputs Inputs, dataSlice []string) string {

	colors := make(map[string]string)

	colors["red"] = "\033[31m"
	colors["green"] = "\033[32m"
	colors["yellow"] = "\033[33m"
	colors["blue"] = "\033[34m"
	colors["cyan"] = "\033[36m"
	colors["reset"] = "\033[0m"

	var builder strings.Builder
	inputSlice := inputs.Str
	subString := inputs.SubString
	colorInput := inputs.Color

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
			for i, ch := range input {

				var ok bool
				var c ColorStruct

				if subString != "" {
					ok, c = CheckColor(input, subString)
				} else {
					builder.WriteString(colors[colorInput])
				}

				if ok && slices.Contains(c.Start, i) {
					builder.WriteString(colors[colorInput])
				}

				//Make sure ch is printable
				if ch >= 32 && ch <= 126 {

					//Calculate formular to get each rows
					print := ((int(ch) - 32) * 9) + row

					//Store string in buffer
					builder.WriteString(dataSlice[print])
				}

				if ok && slices.Contains(c.End, i) {
					builder.WriteString(colors["reset"])
				}

			}
			builder.WriteString("\n")
		}
	}

	return builder.String()

}
