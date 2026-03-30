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

func printString(align string, resultSlice [][]string, width int) {
	terminalWidth, err := terminalWidth()
	if err != nil {
		return
	}

	padding := terminalWidth - width
	fmt.Println(len(resultSlice))

	if align == "right" {
		for _, s := range resultSlice {
			fmt.Print(strings.Repeat(" ", padding) + strings.Join(s, ""))
		}
	} else if align == "center" {
		fmt.Println("Hello")
		for _, s := range resultSlice {
			fmt.Print(strings.Repeat(" ", padding/2) + strings.Join(s, ""))
		}
	} else if align == "justify" {
		fmt.Println()
		for _, s := range resultSlice {
			fmt.Print(strings.Join(s, ""))
		}

	} else {
		for _, s := range resultSlice {
			fmt.Print(strings.Join(s, ""))
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

	// inputSlice := strings.Split(input, "\\n")
	resultSlice, width := operations.AsciiArt(input, data)
	fmt.Println(input.Alignment)
	printString(input.Alignment, resultSlice, width)
	// var builder strings.Builder
	// for _, s := range resultSlice {
	// 	fmt.Print(strings.Repeat(" ", terminalWidth-width) + strings.Join(s, ""))
	// }

	// if input.Alignment == "right" {
	// 	for _, s := range resultSlice {
	// 		builder.WriteString(strings.Repeat(" ", terminalWidth-width) + s)
	// 	}
	// }
	// fmt.Print(builder.String())

	// if input.Output == "" {
	// 	fmt.Print(result)
	// 	// fmt.Println(len(input.Str))
	// 	fmt.Println(width, terminalWidth)
	// 	// fmt.Print(strings.Repeat(" ", terminalWidth-width) + result)
	// } else {
	// 	operations.WriteFile(result, input.Output)
	// }
}
