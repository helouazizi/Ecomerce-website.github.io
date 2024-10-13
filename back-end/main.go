package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("website under structuring")
}

// this function about removing extra spaces
func Expand_Spaces(s string) string {
	// the field work with spaces or more spaces
	result := strings.Fields(s)
	valid_line := strings.TrimSpace(strings.Join(result, " "))

	return valid_line
}

// this function split line by newline
func Split_By_Newline(text string) string {
	result := ""

	for i := 0; i < len(text); i++ {

		char := text[i]
		result += string(char)
		if char == '\n' && i > 1 {
			result += "~ "
		}

	}
	splited_result := strings.Fields(result)
	final_result := strings.Join(splited_result, " ")

	return final_result
}

// this function split line from a )
func Split_line(s string) []string {
	result := []string{}
	startIndex := 0
	endIndex := 0
	delemeter := ')'

	for i, char := range s {
		if char == delemeter {
			endIndex = i
			result = append(result, s[startIndex:endIndex+1])
			startIndex = endIndex + 1
		}
	}
	// append the rest line in the end if exist
	result = append(result, Expand_Spaces(s[startIndex:]))

	return result
}
