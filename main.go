package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Inputfile name missing")
		os.Exit(1)
		return
	}

	inputpath := os.Args[1]
	outputpath := os.Args[2]

	inputfile, err := os.ReadFile(inputpath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
		return
	}

	inputContent := string(inputfile)
	inputContent = ProcessCases(inputContent)
	inputContent = ProcessQuotes(inputContent)
	inputContent = ReplaceSymbol(inputContent)
	inputContent = ProcessPunctuations(inputContent)
	// inputContent = asd(inputContent)

	outputContent := inputContent

	outputFile, err := os.Create(outputpath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.Write([]byte(outputContent))
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		os.Exit(1)
		return
	}
}

func ProcessCases(content string) string {
	hexRegex := regexp.MustCompile(`(\b[\S]+)\s*\(\s*hex\s*\)`)
	binRegex := regexp.MustCompile(`(\b[\S]+)\s*\(\s*bin\s*\)`)
	caseRegex := regexp.MustCompile(`\s*\(\s*(up|cap|low)\s*,*\s*(\d+)?\s*\)`)

	aLowRegex := regexp.MustCompile(`\b(a)(\s+[aoeiuhAOEIUH]\w*)`)
	aUpRegex := regexp.MustCompile(`\b(A)(\s+[AOEIUHaoeiuh]\w*)`)
	anLowRegex := regexp.MustCompile(`\b(a)(?:[Nn])(\s+[^aoeiuhAOEIUH]\w*)`)
	anUpRegex := regexp.MustCompile(`\b(A)(?:[Nn])(\s+[^aoeiuhAOEIUH]\w*)`)

	for i := 0; i < 2; i++ {
		content = aLowRegex.ReplaceAllString(content, "${1}n${2}")
		content = aUpRegex.ReplaceAllString(content, "${1}n${2}")
		content = anLowRegex.ReplaceAllString(content, "${1}${2}")
		content = anUpRegex.ReplaceAllString(content, "${1}${2}")
	}

	content = hexRegex.ReplaceAllStringFunc(content, func(match string) string {
		hexTarget := hexRegex.FindStringSubmatch(match)[1]
		decimal, err := strconv.ParseInt(hexTarget, 16, 64)
		if err != nil {
			return hexTarget
		}
		return strconv.FormatInt(decimal, 10)
	})

	content = binRegex.ReplaceAllStringFunc(content, func(match string) string {
		binTarget := binRegex.FindStringSubmatch(match)[1]

		decimal, err := strconv.ParseInt(binTarget, 2, 64)
		if err != nil {
			return binTarget
		}
		return strconv.FormatInt(decimal, 10)
	})

	for len(caseRegex.FindAllString(content, -1)) > 0 {
		matchIndex := caseRegex.FindAllStringIndex(content, -1)[0]
		startIndex := matchIndex[0]
		endIndex := matchIndex[1]
		wordsTarget := content[:startIndex]
		regCase := caseRegex.FindStringSubmatch(content)[1]
		remainderWords := content[endIndex:]

		quantifierSource := caseRegex.FindStringSubmatch(content)[2]
		quantifier := 1
		if quantifierSource != "" {
			quantifier, _ = strconv.Atoi(quantifierSource)
		}
		if quantifier > 1000 {
			fmt.Println("Quantifiers overflow, quantifier had set to 1000")
			quantifier = 1000
		}
		beforeWord := regexp.MustCompile(`(?:[\w-]+(?:[^\w-]+|\b)){0,` + strconv.Itoa(quantifier) + `}$`)
		res := beforeWord.ReplaceAllStringFunc(wordsTarget, func(match string) string {
			switch regCase {
			case "up":
				return strings.ToUpper(match)
			case "low":
				return strings.ToLower(match)
			case "cap":
				match = strings.ToLower(match)
				return strings.Title(match)
			}
			return match
		})
		content = res + remainderWords
	}

	return strings.TrimSpace(content)
}

func ProcessQuotes(content string) string {
	quotationRegex := regexp.MustCompile(`(\s*)'\s*([^']+)'`)
	singleQuotesRegex := regexp.MustCompile(`(\b\w+)(')(\w+\b)`)
	content = singleQuotesRegex.ReplaceAllString(content, "$1•$3")

	content = quotationRegex.ReplaceAllStringFunc(content, func(match string) string {
		parts := quotationRegex.FindStringSubmatch(match)
		if len(parts) < 3 {
			return match
		}
		return parts[1] + "'" + strings.TrimSpace(parts[2]) + "'"
	})

	return (content)
}

func ReplaceSymbol(input string) string {
	parts := strings.Split(input, "•")
	result := strings.Join(parts, "'")
	return result
}

func ProcessPunctuations(input string) string {
	regsingle := regexp.MustCompile(`(\s*)([,.!?;:]+)`)
	result := regsingle.ReplaceAllString(input, "$2 ")
	regmult := regexp.MustCompile(`([.,!?:;])\s*([.,!?:;])\s*([.,!?:;])`).ReplaceAllString(result, "$1$2$3")
	regtwice := regexp.MustCompile(`([!?])\s*([!?])`).ReplaceAllString(regmult, "$1$2")

	finalResult := strings.ReplaceAll(regtwice, "  ", " ")

	return strings.TrimSpace(finalResult)
}

// func asd(input string) string {
// 	regex := regexp.MustCompile(`([^w])|^'\s*'([^\w])`)
// 	res := regex.ReplaceAllString(input, "$1$2")
// 	fmt.Println(res)
// 	return res
// }
