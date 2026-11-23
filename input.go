package main

import (
	"fmt"
	"unicode"
)

func printStartMenu() bool {
	fmt.Println("Начать новый раунд? в - выход, любой другой ввод - да")
	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		fmt.Println("Ошибка ввода:", err)
		return false
	}
	return !(input == "в")
}

func getCharFromUser() rune {
	for {
		fmt.Println("Введите одну кириллическую букву")
		var input string
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println("Ошибка ввода:", err)
		}
		isCharValid := validateCharInput(input)
		if isCharValid {
			return []rune(input)[0]
		}
		fmt.Println("Некорректный ввод. Попробуйте снова")
	}

}

func validateCharInput(input string) bool {
	runes := []rune(input)
	if len(runes) != 1 {
		return false
	}
	return unicode.Is(unicode.Cyrillic, runes[0])
}
