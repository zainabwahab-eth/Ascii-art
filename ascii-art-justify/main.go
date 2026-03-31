package main

import (
	"ascii-art-justify/operations"
	"fmt"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

const (
	TIOCGWINSZ     = 0x5413
	TIOCGWINSZ_OSX = 1074295912
)

type window struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func terminalWidth() (int, error) {
	w := new(window)
	tio := syscall.TIOCGWINSZ
	if runtime.GOOS == "darwin" {
		tio = TIOCGWINSZ_OSX
	}
	res, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(tio),
		uintptr(unsafe.Pointer(w)),
	)
	if int(res) == -1 {
		return 0, err
	}
	return int(w.Col), nil
}

func printString(align string, resultSlice [][]string, totalWidth int, widthMap map[string]int) {
	terminalWidth, err := terminalWidth()
	if err != nil {
		return
	}

	padding := terminalWidth - totalWidth
	fmt.Println("terminal", terminalWidth)
	fmt.Println("totalwidth", totalWidth)

	switch align {
	case "right":
		for i := 0; i < 8; i++ {
			fmt.Print(strings.Repeat(" ", padding))
			for _, str := range resultSlice {
				fmt.Print(str[i])
			}
			fmt.Println()
		}

	case "centre":
		for i := 0; i < 8; i++ {
			fmt.Print(strings.Repeat(" ", padding/2))
			for _, str := range resultSlice {
				fmt.Print(str[i])
			}
			fmt.Println()
		}

	case "justify":
		if len(resultSlice) == 1 {
			for i := 0; i < 8; i++ {
				for _, str := range resultSlice {
					fmt.Print(str[i])
				}
				fmt.Println()
			}
		} else {
			indPadding := padding / (len(resultSlice) - 1)
			fmt.Println("inpadding", indPadding)
			fmt.Println("padding", padding)
			for i := 0; i < 8; i++ {
				for j, str := range resultSlice {
					if j != len(resultSlice)-1 {
						fmt.Print(str[i] + strings.Repeat(" ", indPadding))
					} else {
						fmt.Print(str[i])
					}
				}
				fmt.Println()
			}
		}
	default:
		for i := 0; i < 8; i++ {
			for _, str := range resultSlice {
				fmt.Print(str[i])
			}
			fmt.Println()
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

	inputSlice := strings.Split(input.Str, " ")

	
	s := []string{}
	resultSlice := [][]string{}
	widthMap := make(map[string]int)
	w := 0
	totalWidth := 0
	for i, word := range inputSlice {
		if i != len(inputSlice)-1 {
			word = word + " "
		}
		s, w = operations.AsciiArt(input, word, data)
		widthMap[word] = w
		totalWidth += w
		resultSlice = append(resultSlice, s)
	}
	// fmt.Println("widthmap", totalWidth)
	// fmt.Print(resultSlice)
	printString(input.Alignment, resultSlice, totalWidth, widthMap)

}
