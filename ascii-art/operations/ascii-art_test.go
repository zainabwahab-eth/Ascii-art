package operations

import (
	"testing"
)

func TestAsciiArt(t *testing.T) {
	type Parameters struct {
		input  [] string
		data []string
	}

	type TestCase struct {
		name     string
		input    Parameters
		expected string
	}


	dataSlice := ReadTextFile("../standard.txt")

	tests := []TestCase{
		{
			name:     "test 1",
			input:    Parameters{
				input: []string{"Hello"},
				data: dataSlice,
			},
			expected: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
`,
		},
		{
			name: "test 2",
			input:    Parameters{
				input: []string{""},
				data: dataSlice,
			},
			expected: `
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AsciiArt(tt.input.input, tt.input.data)
			if result != tt.expected {
				t.Errorf("got \n%q\n, want \n%q", result, tt.expected)
			}
		})
	}
}
