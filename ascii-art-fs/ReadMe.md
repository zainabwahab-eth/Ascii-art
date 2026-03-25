# ASCII Art

A command-line tool written in Go that takes a string and banner file as input and prints it in a graphic representation using ASCII characters based on the banner file entered. Banner file can be standard (default), shadow or thinkertoy. 

This project also works with the previous ascii-art project implemented [Ascii-art-color](./ascii-art-color/ReadMe.md) and [Ascii-art-fs](./ascii-art-fs/ReadMe.md).

## Usage

```bash
go run . "your text here" "banner file here"
```

## Examples

```bash
go run . "Hello" # prints in the default standard banner file
go run . "Hello\nWorld" "shadow" # prints in shawdow graphics with newline between words
go run . "Hello\n\nWorld" "thinkertoy" # blank line between words
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