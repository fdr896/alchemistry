package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// Tool contains tool
type Tool struct {
	tool1, tool2 string
}

func makeTool(s1 string, s2 string) Tool {
	return Tool{tool1: s1, tool2: s2}
}

var thing1, thing2 string

func main() {
	Start()

	for {
		fmt.Scanf("%s %s", &thing1, &thing2)

		thing1 = strings.ToLower(thing1)
		thing2 = strings.ToLower(thing2)

		if thing1 == "показать" && thing2 == "предметы" {
			if err := ShowTools(); err != nil {
				log.Fatal(err)
			}

			Clear()
			continue
		}

		exist1, err := Check(thing1)

		if err != nil {
			log.Fatal(err)
		}

		if !exist1 {
			continue
		}

		exist2, err := Check(thing2)

		if err != nil {
			log.Fatal(err)
		}

		if !exist2 {
			continue
		}

		_, exist3 := tools[makeTool(thing2, thing1)]

		if exist3 {
			thing1, thing2 = thing2, thing1
		}

		newTool, isExist := tools[makeTool(thing1, thing2)]

		if !isExist {
			fmt.Println("Увы ничего не вышло :(")
		} else {
			fmt.Printf("Вы получили: %s\n", newTool)
			AddTool(newTool)
		}

		Clear()
	}
}

// IsToolExist return true if tool is in data.json
func IsToolExist(tool string) (bool, error) {
	file, err := os.Open("data.json")

	if err != nil {
		return false, fmt.Errorf("prolems with opening data.json: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var curTool string
		err := json.Unmarshal([]byte(scanner.Text()), &curTool)

		if err != nil {
			return false, fmt.Errorf("problems with unmarshaling data.json: %v", err)
		}

		if tool == curTool {
			return true, nil
		}
	}

	return false, nil
}

// AddTool adds tool to data.json
func AddTool(tool string) error {

	if exist, _ := IsToolExist(tool); exist {
		return nil
	}

	j, err := json.Marshal(tool)

	if err != nil {
		return fmt.Errorf("problems with marshaling tool: %v", err)
	}

	j = append(j, "\n"...)

	file, _ := os.OpenFile("data.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if _, err := file.Write(j); err != nil {
		return fmt.Errorf("problems with writing tool to file: %v", err)
	}

	return nil
}

// Check checks existens of tool and do right things
func Check(tool string) (bool, error) {
	exist, err := IsToolExist(thing1)

	if err != nil {
		return false, err
	}

	if !exist {
		fmt.Printf("Вы ещё не создали: %s\n", tool)
		Clear()
		return false, nil
	}

	return true, nil
}

// Start starts the game
func Start() {
	os.Create("data.json")
	CallClear()

	AddFormulas()

	AddTool("вода")
	AddTool("земля")
	AddTool("воздух")
	AddTool("огонь")
}

// ShowTools shows all tools from data.json
func ShowTools() error {
	file, err := os.Open("data.json")

	if err != nil {
		return fmt.Errorf("problems with opening file in ShowTools: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var tool string
		err := json.Unmarshal([]byte(scanner.Text()), &tool)

		if err != nil {
			return fmt.Errorf("problems with unmarshaling data.json: %v", err)
		}

		fmt.Println(tool)
	}

	return nil
}
