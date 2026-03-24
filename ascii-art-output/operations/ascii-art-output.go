package operations

import (
	"log"
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
		input = input[:in]
		in = strings.LastIndex(input, subString)
	}
	return true, ColorStruct{
		Start: start,
		End:   end,
	}

}

func ReadTextFile(file string) (error, []string) {
	//Use os.readfile to read our banner file
	data, err := os.ReadFile(file)

	//Return an empty slice if there is an error
	if err != nil {
		return err, []string{}
	}

	cleanStr := strings.ReplaceAll(string(data), "\r", "")

	//Split data by new line
	dataSlice := strings.Split(cleanStr, "\n")

	return nil, dataSlice
}

func GetColorCode(input string) string {
	colorMap := map[string]string{
		"reset":   "\033[0m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"blue":    "\033[34m",
		"yellow":  "\033[33m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"orange":  "\033[38;2;255;165;0m",
	}
	return colorMap[input]
}

func AsciiArt(inputs Inputs, dataSlice []string) string {
	var builder strings.Builder
	inputSlice := inputs.Str
	subString := inputs.SubString
	colorCode := GetColorCode(inputs.Color)
	resetCode := GetColorCode("reset")

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
					builder.WriteString(colorCode)
				}

				if ok && slices.Contains(c.Start, i) {
					builder.WriteString(colorCode)
				}

				//Make sure ch is printable
				if ch >= 32 && ch <= 126 {

					//Calculate formular to get each rows
					print := ((int(ch) - 32) * 9) + row

					//Store string in buffer
					builder.WriteString(dataSlice[print])
				}

				if ok && slices.Contains(c.End, i) {
					builder.WriteString(resetCode)
				}

			}
			builder.WriteString("\n")
		}
		// builder.WriteString(resetCode)
	}

	return builder.String()

}

func WriteFile(data string, submFile string) {
	dataByte := []byte(data)
	err := os.WriteFile(submFile, dataByte, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
