package asciiartoutput

import (
	"fmt"
	"os"
	"strings"
)



func AsciiOutput(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> 'something' ")
		return
	}
	output := args[0]  // Expected to be the output option (e.g., --output=<fileName.txt>)
	input := args[1]   // The input string to generate ASCII art from
	// Case 1: Two arguments, default to standard banner
	if len(args) == 2 {

			filename := strings.TrimPrefix(output, "--output=")
			// Validate filename
			if err := validateFilename(filename); err != nil {
				fmt.Println(err)
				return
			}

			// Default to "standard" banner if no banner is provided
			banner := "standard"

			// Get the appropriate ASCII art file based on the banner
			art, err := GetArtFile(banner)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Generate and write the ASCII art to the output file
			if err := writeAsciiToFile(art, input, filename); err != nil {
				fmt.Println(err)
			}
		return
	}

	// Case 2: Three arguments, handle a specific banner
	if len(args) == 3 {
		banner := args[2] // Banner specified (e.g., standard, shadow, thinkertoy)

		
			filename := strings.TrimPrefix(output, "--output=")

			// Validate filename
			if err := validateFilename(filename); err != nil {
				fmt.Println(err)
				return
			}

			// Get the appropriate ASCII art file based on the banner
			art, err := GetArtFile(banner)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Generate and write the ASCII art to the output file
			if err := writeAsciiToFile(art, input, filename); err != nil {
				fmt.Println(err)
			}
	}
}

// Helper function to validate the output filename
func validateFilename(filename string) error {
	if strings.Contains(filename, "/") {
		return fmt.Errorf("error: file name should not contain '/'")
	}
	if !strings.HasSuffix(filename, ".txt")|| len(filename) <= 4 {
		return fmt.Errorf("error: output file must have a .txt extension")
	}
	return nil
}

// Function to generate and write ASCII art to a file
func writeAsciiToFile(art string, input string, filename string) error {
	// Read the ASCII art file
	asciiArt, err := ReadArtFile(art)
	if err != nil {
		return err
	}

	// Generate the ASCII art based on the input string
	result := GenerateASCIIArt(input, asciiArt)

	// Write the result to the output file
	err = os.WriteFile(filename, []byte(result), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	fmt.Println("Output written to", filename)
	return nil
}