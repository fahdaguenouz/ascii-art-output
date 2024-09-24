package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		if len(args) == 2 || len(args) == 1 || len(args) == 0 {
			fmt.Println("please enter the option and  string to print and the banner 'go run . [OPTION] [STRING] [BANNER]")
		} else {
			fmt.Println("Please enter this format  'go run . [OPTION] [STRING] [BANNER]'")
		}
		return
	}
	output := args[0]
	input := args[1]
	banner := args[2]
	art := ""
	standard := "./art/standard.txt"
	shadow := "./art/shadow.txt"
	thinkertoy := "./art/thinkertoy.txt"
	if strings.HasPrefix(output, "--output=") {
		filename := strings.TrimPrefix(output, "--output=")
		if strings.Contains(filename, "/") {
			fmt.Println("Error: File name should not contain '/'")
			return
		}
		if !strings.HasSuffix(filename, ".txt") {
			fmt.Println("Error: Output file must have a .txt extension")
			return
		}
		if banner == "standard" || banner == "shadow" || banner == "thinkertoy" {
			if banner == "standard" {
				art = standard
			} else if banner == "shadow" {
				art = shadow
			} else if banner == "thinkertoy" {
				art = thinkertoy
			}
			// Open the ASCII art file
			file, err := os.Open(art)
			if err != nil {
				fmt.Println("Error opening the file:", err)
				return
			}
			defer file.Close()
			outputFile, err := os.Create(filename)
			if err != nil {
				fmt.Println("Error creating the output file:", err)
				return
			}
			defer outputFile.Close()

			writer := bufio.NewWriter(outputFile)
			var asciiArt []string

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				asciiArt = append(asciiArt, scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading the file:", err)
				return
			}

			const numLinesPerChar = 8
			startChar := 32
			spaceBetweenChars := 1 // One space between the ASCII characters

			// Each character has an empty line before and after the 8 lines
			linesPerCharacter := numLinesPerChar + spaceBetweenChars // 1 empty line before and after the 8 lines

			lines := strings.Split(input, "\\n")

			for _, line := range lines {
				// Loop over the 8 lines for each character row
				for i := 0; i < numLinesPerChar; i++ {
					// Loop over each character in the line
					for _, char := range line {
						asciiIndex := int(char) - startChar                               // Find the index of the character in the file
						lineIndex := asciiIndex*linesPerCharacter + spaceBetweenChars + i // Account for empty lines

						// Check if the index is valid and print the corresponding line for the character
						if lineIndex < len(asciiArt) {
							writer.WriteString(asciiArt[lineIndex])
						}
						writer.WriteString(" ") // Space between characters
					}
					writer.WriteString("\n") // Newline after each row of characters
				}
				writer.WriteString("\n") // Extra newline after processing each line in the input
			}

			// Flush the buffer to write the content to the file
			writer.Flush()
			fmt.Println("Output written to", filename)
		} else {
			fmt.Println("Invalid banner. Please choose from: standard, shadow, thinkertoy")
			return
		}
	} else {
		fmt.Println("Invalid option. Please use --output=<filename.txt>")

	}
}
