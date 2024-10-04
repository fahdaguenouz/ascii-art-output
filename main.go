package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	
	args := os.Args[1:]

	if len(args) > 3 || len(args) < 2 {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}
	
	output := args[0]
	input := args[1]
	
	art := ""
	standard := "./art/standard.txt"
	shadow := "./art/shadow.txt"
	thinkertoy := "./art/thinkertoy.txt"
	if len(args)==3{
		banner := args[2]
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
						fmt.Println("Please enter a valid character between ascii code 32 and 126")
						return
					}
					index := 9*(int(r)-32) + i 
					result += asciiArt[index]
				}
				result += "\n" // Add newline after finishing the current row of the line
			}
		}
			os.WriteFile(filename,[]byte(result),0644 )
			fmt.Println("Output written to", filename)
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("EX: go run . --output=<fileName.txt> something standard")
			return
		}
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("EX: go run . --output=<fileName.txt> something standard")
	}

	}else if len(args) == 2 {
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
					art = standard
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
							fmt.Println("Please enter a valid character between ascii code 32 and 126")
							return
						}
						index := 9*(int(r)-32) + i 
						result += asciiArt[index]
					}
					result += "\n" // Add newline after finishing the current row of the line
				}
			}
				os.WriteFile(filename,[]byte(result),0644)
	
				fmt.Println("Output written to", filename)
			
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
				fmt.Println("EX: go run . --output=<fileName.txt> something standard")
				return 
		}
	}
}
