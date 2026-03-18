package operations

import (
	"testing"
)

func TestAsciiArt(t *testing.T) {

	type Parameters struct {
		input Inputs
		data  []string
	}

	type TestCase struct {
		name     string
		inputs   Parameters
		expected string
	}

	dataSlice := ReadTextFile("../standard.txt")

	tests := []TestCase{
		{
			name: "test 1",
			inputs: Parameters{
				input: Inputs{
					Str:       []string{"Hello"},
					SubString: "l",
					Color:     "red",
				},
				data: dataSlice,
			},
			expected: ` _    _         x1b[31m_  x1b[0mx1b[31m_  x1b[0m         
| |  | |        \x1b[31m| | \x1b[0m\x1b[31m| | \x1b[0m         
| |__| |   ___  \x1b[31m| | \x1b[0m\x1b[31m| | \x1b[0m  ___   
|  __  |  / _ \ \x1b[31m| | \x1b[0m\x1b[31m| | \x1b[0m / _ \  
| |  | | |  __/ \x1b[31m| | \x1b[0m\x1b[31m| | \x1b[0m| (_) | 
|_|  |_|  \___| \x1b[31m|_| \x1b[0m\x1b[31m|_| \x1b[0m \___/  
                                
                                
`,

			//" _    _         \x1b[31m _  \x1b[0m\x1b[31m _  \x1b[0m        \n| |  | |        \x1b[31m| | \x1b[0m\x1b[31m| | \x1b[0m        \n| |__| |   ___  \x1b[31m| | \x1b[0m\x1b[31m| | \x1b[0m  ___   \n|  __  |  / _ \\ \x1b[31m| | \x1b[0m\x1b[31m| | \x1b[0m / _ \\  \n| |  | | |  __/ \x1b[31m| | \x1b[0m\x1b[31m| | \x1b[0m| (_) | \n|_|  |_|  \\___| \x1b[31m|_| \x1b[0m\x1b[31m|_| \x1b[0m \\___/  \n                \x1b[31m    \x1b[0m\x1b[31m    \x1b[0m        \n                \x1b[31m    \x1b[0m\x1b[31m    \x1b[0m        \n"
		},
// 		{
// 			name: "test 2",
// 			inputs: Parameters{
// 				input: Inputs{
// 					Str:       []string{""},
// 					SubString: "",
// 					Color:     "yellow",
// 				},
// 				data: dataSlice,
// 			},
// 			expected: `
// `,
// 		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AsciiArt(tt.inputs.input, tt.inputs.data)
			if result != tt.expected {
				t.Errorf("got \n%q\n, want \n%q", result, tt.expected)
			}
		})
	}
}
