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

			var result string
		lines := strings.Split(input, "\\n")

		for _, line := range lines {
			if line == "" {
				result += "\n"
				continue
			}
			// Iterate over each row of the ASCII art (0 to 7, for the 8 rows)
			for i := 1; i <= 8; i++ {
				for _, r := range line {
					// Ensure the character is within the valid ASCII range
					if r < 32 || r > 126 {
						fmt.Println("Please enter a valide character between ascii code 32 and 126")
						return
					}
					index := 9*(int(r)-32) + i 
					result += asciiArt[index]
				}
				result += "\n" // Add newline after finishing the current row of the line
			}
		}
			writer.WriteString(result)

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
