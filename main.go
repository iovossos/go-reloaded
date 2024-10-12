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

func readfile(input string, output string) {
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	transformedData := transformText(string(data))
	writefile(output, transformedData)
}

func transformText(dataStr string) string {
	lines := strings.Split(dataStr, "\n") // Split by lines
	for i, line := range lines {
		lines[i] = transformHex(line)
		lines[i] = transformBin(lines[i])
		lines[i] = transformCase(lines[i])
		lines[i] = transformPunctuation(lines[i])
		lines[i] = transformArticles(lines[i])
	}
	return strings.Join(lines, "\n") // Join lines with newlines to preserve formatting
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
	upRegex := regexp.MustCompile(`(\b(?:\w+\b\s*){1,})(?:\(up(?:,\s*(\d+))?\))`)
	dataStr = upRegex.ReplaceAllStringFunc(dataStr, func(match string) string {
		words := strings.Fields(match[:strings.LastIndex(match, " (up")]) // Extract words before (up)
		if strings.Contains(match, ",") {                                 // Check if there's a number after (up,
			num := match[strings.LastIndex(match, ",")+1 : strings.LastIndex(match, ")")]
			n, _ := strconv.Atoi(strings.TrimSpace(num))
			for i := len(words) - n; i < len(words); i++ {
				words[i] = strings.ToUpper(words[i])
			}
		} else {
			words[len(words)-1] = strings.ToUpper(words[len(words)-1]) // Only the previous word
		}
		return strings.Join(words, " ")
	})

	// Transform (low)
	lowRegex := regexp.MustCompile(`(\b(?:\w+\b\s*){1,})(?:\(low(?:,\s*(\d+))?\))`)
	dataStr = lowRegex.ReplaceAllStringFunc(dataStr, func(match string) string {
		words := strings.Fields(match[:strings.LastIndex(match, " (low")])
		if strings.Contains(match, ",") {
			num := match[strings.LastIndex(match, ",")+1 : strings.LastIndex(match, ")")]
			n, _ := strconv.Atoi(strings.TrimSpace(num))
			for i := len(words) - n; i < len(words); i++ {
				words[i] = strings.ToLower(words[i])
			}
		} else {
			words[len(words)-1] = strings.ToLower(words[len(words)-1])
		}
		return strings.Join(words, " ")
	})

	// Transform (cap)
	capRegex := regexp.MustCompile(`(\b(?:\w+\b\s*){1,})(?:\(cap(?:,\s*(\d+))?\))`)
	dataStr = capRegex.ReplaceAllStringFunc(dataStr, func(match string) string {
		words := strings.Fields(match[:strings.LastIndex(match, " (cap")])
		if strings.Contains(match, ",") {
			num := match[strings.LastIndex(match, ",")+1 : strings.LastIndex(match, ")")]
			n, _ := strconv.Atoi(strings.TrimSpace(num))
			for i := len(words) - n; i < len(words); i++ {
				words[i] = strings.Title(words[i])
			}
		} else {
			words[len(words)-1] = strings.Title(words[len(words)-1])
		}
		return strings.Join(words, " ")
	})

	return dataStr
}

func transformPunctuation(dataStr string) string {
	// Handle group punctuation (e.g., "...", "!?") without splitting them
	groupPunctuationRegex := regexp.MustCompile(`\s*([.]{3}|[!?]{2,3})\s*`)
	dataStr = groupPunctuationRegex.ReplaceAllString(dataStr, "$1") // Don't add spaces inside punctuation groups

	// Handle standard punctuation like ., ,, !, ?, : and ; (close to the previous word, and ensure a space after)
	punctuationRegex := regexp.MustCompile(`\s*([.,!?;:])\s*`)
	dataStr = punctuationRegex.ReplaceAllString(dataStr, "$1 ") // Ensure space after punctuation

	// Handle single quotes surrounding words and ensure they close properly before punctuation
	quoteRegex := regexp.MustCompile(`'\s*(.*?)\s*'`)
	dataStr = quoteRegex.ReplaceAllStringFunc(dataStr, func(match string) string {
		content := strings.TrimSpace(match[1 : len(match)-1])
		// Ensure quotes are followed by a space only if they aren't directly followed by punctuation
		if len(content) > 0 && strings.ContainsAny(content[len(content)-1:], ".,!?;:") {
			return "'" + content[:len(content)-1] + "' " + content[len(content)-1:]
		}
		return "'" + content + "'"
	})

	// Fix any extra spaces caused by punctuation handling (e.g., spaces before commas or periods)
	dataStr = strings.ReplaceAll(dataStr, " ,", ",")
	dataStr = strings.ReplaceAll(dataStr, " .", ".")
	dataStr = strings.ReplaceAll(dataStr, " !", "!")
	dataStr = strings.ReplaceAll(dataStr, " ?", "?")
	dataStr = strings.ReplaceAll(dataStr, " ;", ";")
	dataStr = strings.ReplaceAll(dataStr, " :", ":")

	// Trim any extra spaces from the final result
	return strings.TrimSpace(dataStr)
}

func transformArticles(dataStr string) string {
	// Replace "a" with "an" before a vowel or 'h', and handle uppercase words too
	articleRegex := regexp.MustCompile(`\b([Aa])\s+([aeiouhAEIOUH])`)
	return articleRegex.ReplaceAllStringFunc(dataStr, func(match string) string {
		if match[0] == 'A' {
			return "An " + match[2:]
		}
		return "an " + match[2:]
	})
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
	readfile(inname, outname)
}
