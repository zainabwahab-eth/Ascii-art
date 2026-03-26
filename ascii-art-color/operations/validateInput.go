package operations

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Inputs struct {
	Str       []string
	SubString string
	Color     string
}

func isValidColor(color string) bool {
	if strings.HasPrefix(color, "#") && len(color) == 7 {
		return true
	}
	if strings.HasPrefix(color, "rgb(") && strings.HasSuffix(color, ")") {
		return true
	}
	namedColors := []string{"red", "blue", "yellow", "green", "magenta", "cyan", "orange"}
	return slices.Contains(namedColors, color)
}

func ValidateInput() (bool, Inputs) {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--color=") {
			fmt.Println(`Usage: go run . [OPTION] [STRING] \n\nEX: go run . --color=<color> <substring to be colored> "something"`)
			os.Exit(1)
		}
	}
	color := flag.String("color", "", "color to use")

	flag.Parse()

	args := flag.Args()

	//if color is specified check inputs
	if (*color != "") && (len(args) != 1 && len(args) != 2) {
		fmt.Println(`Usage: go run . [OPTION] [STRING] \n\nEX: go run . --color=<color> <substring to be colored> "something"`)
		return false, Inputs{}
	}

	//if color is not specified check inputs
	if *color == "" && len(args) != 1 {
		fmt.Println(`Usage: go run . [OPTION] [STRING] \n\nEX: go run . --color=<color> <substring to be colored> "something"`)
		return false, Inputs{}
	}

	if *color != "" && !isValidColor(strings.ToLower(*color)) {
		fmt.Println("Invalid color format. Use named colors, #hex, or rgb(r,g,b)")
		return false, Inputs{}
	}

	str := ""
	subString := ""

	if len(args) < 2 {
		str = args[len(args)-1]
	} else if len(args) == 2 {
		str = args[len(args)-1]
		subString = args[0]
	}

	strSlice := strings.Split(str, "\\n")

	return true, Inputs{
		Str:       strSlice,
		SubString: subString,
		Color:     *color,
	}
}
