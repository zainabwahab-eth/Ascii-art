# ASCII Art Color

A command-line tool written in Go that takes a color flag, substring and string as input and prints the string in a graphic representation using ASCII characters with the substring colored in the specified color. If no substring is specified the entire string is colored in the specified color.

## Usage

```bash
go run . --color="your specified color" "your substring to be colored" "your text here"
```

## Examples

```bash
go run . --color="red" "world" "Helloworld"
go run . --color="yellow" "Hello\nWorld" # no substring specified so all text will be colored yellow
go run . "" # empty string
```

## Features

- Supports uppercase and lowercase letters
- Supports uppercase and lowercase colors
- Supports the following colors; red, green, yellow, cyan, blue
- Supports numbers and special characters
- Supports spaces
- Supports \n for newlines in input

## Requirements

- Go 1.18+
- Only standard Go packages are used