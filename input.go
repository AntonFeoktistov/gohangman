package main

import "fmt"

func printStartMenu() bool {
	fmt.Println("Начать новый раунд? в - выход, любой другой ввод - да")
	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		fmt.Println("Ошибка ввода:", err)
		return false
	}
	return input == "в"
}
