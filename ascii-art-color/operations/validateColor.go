package operations

import (
	"fmt"
	"strconv"
	"strings"
)

func GetColorCode(input string) string {
    if strings.HasPrefix(input, "#") {
        return hexToAnsi(input)
    }
    if strings.HasPrefix(input, "rgb(") {
        return rgbToAnsi(input)
    }
    // fallback to named colors
    return namedToAnsi(input)
}

func namedToAnsi(input string) string {
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

func hexToAnsi(hex string) string {
	//#FF5733
	trimmedHex := strings.TrimPrefix(hex, "#")

	//convert from hexadecimal to decimal
	r, _ := strconv.ParseInt(trimmedHex[0:2], 16, 64) //FF
	g, _ := strconv.ParseInt(trimmedHex[2:4], 16, 64) //57
	b, _ := strconv.ParseInt(trimmedHex[4:6], 16, 64) //33

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)

}

func rgbToAnsi(rgb string) string {
	//rgb(255, 87, 51)
	trimmedRgb := strings.TrimPrefix(rgb, "rgb(")
	trimmedRgb = strings.TrimSuffix(trimmedRgb, ")")

	//Split by , to get each value
	rgbSlice := strings.Split(trimmedRgb, ",")

	//Trim space and Convert each value to int
	r, _ := strconv.Atoi(strings.TrimSpace(rgbSlice[0]))
    g, _ := strconv.Atoi(strings.TrimSpace(rgbSlice[1]))
    b, _ := strconv.Atoi(strings.TrimSpace(rgbSlice[2]))
	
	fmt.Println(r, g, b)
    return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)


}
