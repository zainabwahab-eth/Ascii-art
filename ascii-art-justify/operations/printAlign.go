package operations

import (
	"fmt"
	"os"
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

func TerminalWidth() (int, error) {
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

func PerformAscii(str string, input Inputs, dataSlice []string) ([][]string, int) {
	inputSlice := strings.Split(str, " ")

	s := []string{}
	resultSlice := [][]string{}
	// widthMap := make(map[string]int)
	w := 0
	totalWidth := 0

	for i, word := range inputSlice {
		if i != len(inputSlice)-1 {
			word = word + " "
		}
		s, w = AsciiArt(input, word, dataSlice)
		// widthMap[word] = w
		totalWidth += w
		resultSlice = append(resultSlice, s)
	}

	return resultSlice, totalWidth
}

func PrintRight(resultSlice [][]string, padding int) {
	for i := 0; i < 8; i++ {
		fmt.Print(strings.Repeat(" ", padding))
		for _, str := range resultSlice {
			fmt.Print(str[i])
		}
		fmt.Println()
	}
}

func PrintCentre(resultSlice [][]string, padding int) {
	for i := 0; i < 8; i++ {
		fmt.Print(strings.Repeat(" ", padding/2))
		for _, str := range resultSlice {
			fmt.Print(str[i])
		}
		fmt.Println()
	}
}

func PrintJustify(resultSlice [][]string, padding int) {
	if len(resultSlice) == 1 {
		for i := 0; i < 8; i++ {
			for _, str := range resultSlice {
				fmt.Print(str[i])
			}
			fmt.Println()
		}
	} else {
		indPadding := padding / (len(resultSlice) - 1)
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
}

func PrintDefault(resultSlice [][]string) string {
	var res strings.Builder
	for i := 0; i < 8; i++ {
		for _, str := range resultSlice {
			res.WriteString(str[i])
		}
		res.WriteString("\n")
	}

	return res.String()
}

func WriteFile(data string, submFile string) {
	f, err := os.OpenFile(submFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer f.Close()
	f.WriteString(data)

}
