package config

import (
	"bufio"
	"fmt"
	"os"
)

func LoadBeauty() {
	args := os.Args
	if (len(args) > 1) && (args[1] == "beauty") {
		showWelcomeMessage()
	}
}

func showWelcomeMessage() {
	file, err := os.Open("welcome.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("error opening welcome file: %v", err))
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("error closing welcome file: %v", err))
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
