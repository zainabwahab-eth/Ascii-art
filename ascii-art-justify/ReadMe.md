# ASCII Art Justify

A command-line tool written in Go that takes an align flag, string and banner file as input and writes the string in a graphic representation using ASCII characters to the output file entered (eg example.txt). The string is written based on the alignment type entered (left, right, center and justify)

Banner file can be either standard, shadow or thinkertoy. If no banner file is specified the string is written based on the standard banner file.

This project also works with the previous ascii-art project implemented [Ascii-art-color](./ascii-art-color/ReadMe.md), [Ascii-art-fs](./ascii-art-fs/ReadMe.md) and [Ascii-art-output](./ascii-art-output/ReadMe.md)

## Usage

```bash
go run . --align="your alignment type" "your string here" "your banner file"
```

## Examples

```bash
go run . --align="center" "Hello World" "shadow"
go run . --align="justify" "Hello hi\nWorld" "thinkertoy" #works with \n
go run . "" # empty string
```

## Features

- Supports uppercase and lowercase letters
- Supports alignment right, left, center and justify
- Supports "\n" (new line)
- Supports numbers and special characters
- Supports spaces
- Supports \n for newlines in input

## Requirements

- Go 1.18+
- Only standard Go packages are used