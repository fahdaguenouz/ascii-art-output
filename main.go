package main

import (
	
	"fmt"
	"os"
	"strings"
	"asciiartoutput/functions"
)

func main() {
	args := os.Args[1:]


if len(args)==1{
	
	input:=args[0]
	banner:="standard"
	result:=asciiartoutput.AsciiNormal(input,banner)
	fmt.Print(result)
	return
}else if len(args)==2{
	 if strings.HasPrefix(args[0], "--output=") {
		asciiartoutput.AsciiOutput(args)
	}else if args[1]=="standard"||args[1]=="thinkertoy"||args[1]=="shadow"{
	input:=args[0]
    banner:=args[1]
    result:=asciiartoutput.AsciiFs(input,banner)
    fmt.Print(result)
    return
}else {
	fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --output=<fileName.txt> something standard")
    return
}

}else if len(args)==3{

	if strings.HasPrefix(args[0], "--output=") {
		asciiartoutput.AsciiOutput(args)
	}else {
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

} else {
	fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --output=<fileName.txt> something standard")

		return
	}
}