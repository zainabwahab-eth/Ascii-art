package operations

import (
	"strings"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	dataSlice := ReadTextFile("../standard.txt")

	tests := []struct {
		name      string
		input     Inputs
		substring string
	}{
		{
			name: "test 1",
			input: Inputs{
				Str:       []string{"Hello"},
				SubString: "l",
				Color:     "red",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AsciiArt(tt.input, dataSlice)
			if result == "" {
				t.Errorf("expected output but got empty string")
			}
			// check color code is present
			if !strings.Contains(result, "\033[31m") {
				t.Errorf("expected red color code in output")
			}
			// check reset code is present
			if !strings.Contains(result, "\033[0m") {
				t.Errorf("expected reset code in output")
			}
		})
	}
}
