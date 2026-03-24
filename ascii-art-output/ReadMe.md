# ASCII Art Output

A command-line tool written in Go that takes an output flag, string and banner file as input and writes the string in a graphic representation using ASCII characters to the output file entered (eg example.txt). The string is written based on the banner file entered.

Banner file can be either standard, shadow or thinkertoy. If no banner file is specified the string is written based on the standard banner file.

This project also works with the previous ascii-art project implemented [Ascii-art-color](./ascii-art-color/ReadMe.md).

## Usage

```bash
go run . --output="your specified output file" "your string here" "your banner file"
```

## Examples

```bash
go run . --output="test.txt" "world" "shadow"
go run . --output="example.txt" "Hello\nWorld" # no banner file specified so string is written in example.txt based on the standard banner file.
go run . "" # empty string
```

## Features

- Supports uppercase and lowercase letters
- Supports numbers and special characters
- Supports spaces
- Supports \n for newlines in input

## Requirements

- Go 1.18+
- Only standard Go packages are used