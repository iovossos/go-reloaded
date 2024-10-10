
# Text Transformer

This is a simple Go project that reads an input text file, applies multiple transformations based on specific rules, and writes the transformed content to an output text file.

## Features
- **Hexadecimal to Decimal Conversion**: Converts any hexadecimal number followed by `(hex)` into its decimal equivalent.
- **Binary to Decimal Conversion**: Converts any binary number followed by `(bin)` into its decimal equivalent.
- **Case Transformations**: 
  - `(up)` transforms the previous word into uppercase.
  - `(low)` transforms the previous word into lowercase.
  - `(cap)` transforms the previous word to capitalized.
  - If a number is specified, like `(up, 2)`, it will transform that many previous words.
- **Punctuation Handling**: Ensures punctuation like `.,!?;:` are correctly formatted with proper spacing.
- **Article Correction**: Automatically converts "a" to "an" if the following word starts with a vowel or 'h'.
- **Quote Handling**: Ensures single quotes `'` are correctly placed around the quoted text, and properly closed before punctuation.

## Usage
To run the program, use the following command:

```bash
go run main.go <input_file> <output_file>
```

- `<input_file>`: Path to the input text file.
- `<output_file>`: Path where the transformed output will be written.

Example:

```bash
go run main.go input.txt output.txt
```

## Requirements
- Go 1.16 or later

## Installation
1. Clone this repository:
   ```bash
   git clone https://platform.zone01.gr/git/ivossos/go-reloaded
   ```
2. Navigate to the project directory:
   ```bash
   cd go-reloaded
   ```

3. Run the program:
   ```bash
   go run main.go input.txt output.txt
   ```
