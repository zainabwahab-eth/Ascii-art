package operations

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Inputs struct {
	Str       string
	SubString string
	Color     string
	Output    string
	Banner    string
	Alignment string
}

func PrintColorError() {
	fmt.Println(`Usage: go run . [OPTION] [STRING] \n\nEX: go run . --color=<color> <substring to be colored> "something"`)

}

func PrintOutputError() {
	fmt.Println(`Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard`)

}

func PrintAlignError() {
	fmt.Println(`Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=<alignment> something standard`)

}

func ParseNoFlag() Inputs {
	args := flag.Args()

	if len(args) > 1 {
		return Inputs{}
	}

	str := args[0]

	// strSlice := strings.Split(str, "\\n")

	return Inputs{
		Str: str,
	}
}

func ParseColorFlag(color string) Inputs {
	args := flag.Args()

	//if color is specified check inputs
	if (color != "") && (len(args) != 1 && len(args) != 2) {
		PrintColorError()
		return Inputs{}
	}

	acceptedColors := []string{"red", "blue", "yellow", "green", "magenta", "cyan", "orange"}

	if color != "" && !slices.Contains(acceptedColors, strings.ToLower(color)) {
		fmt.Println(color + ` not available, please try red, blue, yellow, green, magenta, cyan, orange`)
		return Inputs{}
	}

	str := ""
	subString := ""

	str = args[0]

	if len(args) == 2 {
		str = args[1]
		subString = args[0]
	}

	// strSlice := strings.Split(str, "\\n")

	return Inputs{
		Str:       str,
		SubString: subString,
		Color:     strings.ToLower(color),
	}
}

func ParseOutputFlag(output string) Inputs {
	args := flag.Args()

	if (output != "") && (len(args) != 1 && len(args) != 2) {
		PrintOutputError()
		return Inputs{}
	}

	str := args[0]
	banner := "standard"
	if len(args) > 1 {
		banner = args[1]
	}

	// strSlice := strings.Split(str, "\\n")

	return Inputs{
		Str:    str,
		Banner: banner,
		Output: output,
	}
}

func ParseAlignFlag(align string) Inputs {
	args := flag.Args()

	if (align != "") && (len(args) != 1 && len(args) != 2) {
		PrintAlignError()
		return Inputs{}
	}

	acceptedAlignment := []string{"left", "center", "right", "justify"}

	if align != "" && !slices.Contains(acceptedAlignment, strings.ToLower(align)) {
		fmt.Println(align + ` not available, please try left, right, center, justify`)
		return Inputs{}
	}

	str := args[0]
	banner := "standard"
	if len(args) > 1 {
		banner = args[1]
	}

	// strSlice := strings.Split(str, "\\n")

	return Inputs{
		Str:       str,
		Banner:    banner,
		Alignment: strings.ToLower(align),
	}
}

func ParseColorAndOutputFlag(color string, output string) Inputs {
	args := flag.Args()

	//if color is specified check inputs
	if len(args) < 1 || len(args) >= 4 {
		fmt.Println(`Usage: go run . [OPTION] [STRING] \n\nEX: go run . --color=<color> <substring to be colored> "something"`)
		return Inputs{}
	}

	acceptedColors := []string{"red", "blue", "yellow", "green", "magenta", "cyan", "orange"}

	if color != "" && !slices.Contains(acceptedColors, strings.ToLower(color)) {
		fmt.Println(color + ` not available, please try red, blue, yellow, green, magenta, cyan, orange`)
		return Inputs{}
	}

	str := ""
	subString := ""
	banner := "standard"

	str = args[0]

	if len(args) == 2 {
		str = args[1]
		subString = args[0]
	}

	if len(args) > 2 {
		str = args[1]
		subString = args[0]
		banner = args[2]
	}

	// strSlice := strings.Split(str, "\\n")

	return Inputs{
		Str:       str,
		SubString: subString,
		Color:     strings.ToLower(color),
		Output:    output,
		Banner:    banner,
	}
}

func ValidateInput() (bool, Inputs) {
	acceptedFlags := []string{"--color=", "--output=", "--align="}

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && !slices.Contains(acceptedFlags, strings.Split(arg, "=")[0]+"=") {
			if strings.Contains(arg, "color") {
				PrintColorError()
			} else if strings.Contains(arg, "output") {
				PrintOutputError()
			} else {
				PrintAlignError()
			}
			os.Exit(1)
		}
	}

	inputs := Inputs{}
	output := flag.String("output", "", "Output file")
	color := flag.String("color", "", "color to use")
	align := flag.String("align", "", "Text alignment")
	flag.Parse()

	switch {
	case *color == "" && *output == "" && *align == "":
		inputs = ParseNoFlag()

	// case *output != "" && *color != "":
	// 	inputs = ParseColorAndOutputFlag(*color, *output)

	case *color != "":
		inputs = ParseColorFlag(*color)

	case *output != "":
		inputs = ParseOutputFlag(*output)

	case *align != "":
		inputs = ParseAlignFlag(*align)
	}

	return true, inputs
}
