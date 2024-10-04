package asciiartoutput

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to get the appropriate art file based on the banner type
func GetArtFile(banner string) (string, error) {
	standard := "./art/standard.txt"
	shadow := "./art/shadow.txt"
	thinkertoy := "./art/thinkertoy.txt"

	switch banner {
	case "standard":
		return standard, nil
	case "shadow":
		return shadow, nil
	case "thinkertoy":
		return thinkertoy, nil
	default:
		return "", fmt.Errorf("invalid banner. Please choose from: standard, shadow, thinkertoy\n ")
	}
}

func ReadArtFile(art string) ([]string, error) {
	file, err := os.Open(art)
	if err != nil {
		return nil, fmt.Errorf("error opening the file: %v\n ", err)
	}
	defer file.Close()

	var asciiArt []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		asciiArt = append(asciiArt, scanner.Text())
	}
	if len(asciiArt) != 855 {
		return nil, fmt.Errorf("file error: expected 855 lines, but got %d lines", len(asciiArt))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading the file: %v\n ", err)
	}

	return asciiArt, nil
}

// Function to generate ASCII art from the input string
func GenerateASCIIArt(input string, asciiArt []string) string {
	var result string
	lines := strings.Split(input, "\\n")
count:=0
	for _, line := range lines {
		if line == "" {
			result += "\n"
			count++
			continue
		}
		for i := 1; i <= 8; i++ {
			for _, r := range line {
				if r < 32 || r > 126 {
					fmt.Println("invalid character: please enter a valid character between ASCII code 32 and 126")
					return ""
				}
				index := 9*(int(r)-32) + i
				result += asciiArt[index]
			}
			result += "\n"

		}
	}
	if count == len(result) {
		result = result[:len(result)-1]
	}

	return result
}
