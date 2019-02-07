package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tools map[Tool]string = map[Tool]string{
	makeTool("вода", "вода"): "море",
}

// Parser return slice of strings
func Parser(str string) (res [3]string) {
	var curRes string
	space := " "

	for _, val := range str {
		val := string(val)
		if val != space && val != "=" && val != "+" {
			curRes += string(val)
		} else if val == space {
			continue
		} else if val == "=" {
			curRes = strings.ToLower(curRes)
			res[0] = curRes
			curRes = ""
		} else if val == "+" {
			curRes = strings.ToLower(curRes)
			res[1] = curRes
			curRes = ""
		}
	}

	res[2] = strings.ToLower(curRes)

	return
}

// AppendToTools adds formula to tools
func AppendToTools(formula string) {
	var res, tool1, tool2 string

	tmp := Parser(formula)

	res, tool1, tool2 = tmp[0], tmp[1], tmp[2]

	tools[makeTool(tool1, tool2)] = res
}

// AddFormulas adds every formulas to map
func AddFormulas() error {
	file, err := os.Open("formulas.txt")

	if err != nil {
		return fmt.Errorf("problems with opening file in AddFormulas %v", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		formula := string(scanner.Text())

		AppendToTools(formula)
	}

	return nil
}
