package main

import (
	"fmt"
	"slices"
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

func getCharFromUser(set []rune) (rune, []rune) {
	for {
		fmt.Println("Введите одну кириллическую букву")
		var input string
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println("Ошибка ввода:", err)
		}
		isCharValid := validateCharInput(input, set)
		if !isCharValid {
			fmt.Println("Некорректный ввод. Попробуйте снова")
			continue
		}
		char := []rune(input)[0]
		set, isInSet := checkCharAndAdd(char, set)
		if isInSet {
			fmt.Println("Вы уже вводили эту букву. Введите другую")
			continue
		}
		return char, set
	}

}

func validateCharInput(input string, set []rune) bool {
	runes := []rune(input)
	if len(runes) != 1 {
		return false
	}
	return unicode.Is(unicode.Cyrillic, runes[0])
}

func checkCharAndAdd(char rune, set []rune) ([]rune, bool) {
	if slices.Contains(set, char) {
		return set, true
	}
	return append(set, char), false
}
