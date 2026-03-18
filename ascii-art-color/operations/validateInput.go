package operations

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Inputs struct {
	Str       []string
	SubString string
	Color     string
}

func ValidateInput() (bool, Inputs) {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--color=") {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			os.Exit(1)
		}
	}
	color := flag.String("color", "", "color to use")

	flag.Parse()

	args := flag.Args()

	//if color is specified check inputs
	if (*color != "") && (len(args) != 1 && len(args) != 2) {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		return false, Inputs{}
	}

	//if color is not specified check inputs
	if *color == "" && len(args) != 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING]2")
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
		Color:     strings.ToLower(*color),
	}
}
