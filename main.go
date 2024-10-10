package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func createfile(output string) {
	file, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Printf("Output created at file: %s\n", file.Name())
}

func readfile(input string) {
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	transformedData := transformText(string(data))
	writefile("output.txt", transformedData)
}

func transformText(dataStr string) string {
	dataStr = transformHex(dataStr)
	dataStr = transformBin(dataStr)
	dataStr = transformCase(dataStr)
	dataStr = transformPunctuation(dataStr)
	dataStr = transformArticles(dataStr)
	return dataStr
}

func transformHex(dataStr string) string {
	hexRegex := regexp.MustCompile(`(\b[0-9A-Fa-f]+\b) \(hex\)`)
	return hexRegex.ReplaceAllStringFunc(dataStr, func(match string) string {
		hexNum := match[:len(match)-6] // Remove " (hex)"
		decimal, err := strconv.ParseInt(hexNum, 16, 0)
		if err != nil {
			return match // Return original if there's an error
		}
		return fmt.Sprintf("%d", decimal)
	})
}

func transformBin(dataStr string) string {
	binRegex := regexp.MustCompile(`(\b[01]+\b) \(bin\)`)
	return binRegex.ReplaceAllStringFunc(dataStr, func(match string) string {
		binNum := match[:len(match)-6] // Remove " (bin)"
		decimal, err := strconv.ParseInt(binNum, 2, 0)
		if err != nil {
			return match // Return original if there's an error
		}
		return fmt.Sprintf("%d", decimal)
	})
}

func transformCase(dataStr string) string {
	// Transform (up)
	upRegex := regexp.MustCompile(`(\b\w+\b) \(up(?:, \d+)?\)`)
	dataStr = upRegex.ReplaceAllStringFunc(dataStr, func(match string) string {
		words := strings.Fields(match[:len(match)-6]) // Remove " (up)"
		if strings.Contains(match, ",") {
			numWords := strings.TrimSpace(match[len(match)-5:])
			n, _ := strconv.Atoi(numWords)
			for i := 0; i < n && i < len(words); i++ {
				words[i] = strings.ToUpper(words[i])
			}
		} else {
			words[0] = strings.ToUpper(words[0])
		}
		return strings.Join(words, " ")
	})

	// Implement transformations for (low) and (cap) similarly...
	// For example:
	lowRegex := regexp.MustCompile(`(\b\w+\b) \(low(?:, \d+)?\)`)
	dataStr = lowRegex.ReplaceAllStringFunc(dataStr, func(match string) string {
		words := strings.Fields(match[:len(match)-6]) // Remove " (low)"
		if strings.Contains(match, ",") {
			numWords := strings.TrimSpace(match[len(match)-5:])
			n, _ := strconv.Atoi(numWords)
			for i := 0; i < n && i < len(words); i++ {
				words[i] = strings.ToLower(words[i])
			}
		} else {
			words[0] = strings.ToLower(words[0])
		}
		return strings.Join(words, " ")
	})

	// For (cap)
	capRegex := regexp.MustCompile(`(\b\w+\b) \(cap(?:, \d+)?\)`)
	dataStr = capRegex.ReplaceAllStringFunc(dataStr, func(match string) string {
		words := strings.Fields(match[:len(match)-6]) // Remove " (cap)"
		if strings.Contains(match, ",") {
			numWords := strings.TrimSpace(match[len(match)-5:])
			n, _ := strconv.Atoi(numWords)
			for i := 0; i < n && i < len(words); i++ {
				words[i] = strings.Title(words[i])
			}
		} else {
			words[0] = strings.Title(words[0])
		}
		return strings.Join(words, " ")
	})

	return dataStr
}

func transformPunctuation(dataStr string) string {
	// Implement punctuation transformation logic here
	return dataStr
}

func transformArticles(dataStr string) string {
	// Implement article transformation logic here
	return dataStr
}

func writefile(output string, data string) {
	err := os.WriteFile(output, []byte(data), 0o644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transformed data written to:", output)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}
	inname := os.Args[1]
	outname := os.Args[2]
	createfile(outname)
	readfile(inname)
}
