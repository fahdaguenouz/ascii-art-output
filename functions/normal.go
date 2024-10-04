package asciiartoutput

import (
	"fmt"
)

func AsciiNormal(input string, banner string) string {

	artFile, err := GetArtFile(banner)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	asciiArt, err := ReadArtFile(artFile)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	
	result := GenerateASCIIArt(input, asciiArt)
	return result
}