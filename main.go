package main

import (
	"fmt"
	"log"
	"os"
)

const maxMistakes int = 5

func main() {
	dictionary, err := readWordsFromFile("words.txt")
	if err != nil {
		log.Fatalf("не удалось прочитать файл: %v", err)
	}

	for {
		startGame := printStartMenu()
		if !startGame {
			os.Exit(0)
		}

		mistakes := maxMistakes
		word := getRandomWord(dictionary)
		mask := makeMask(word)
		guessedChars := make([]rune, 0)

		fmt.Println("Слово загадано. Букв:", len([]rune(word)))
		fmt.Println(word)
		for !isGameFinished(word, mask, mistakes) {
			printMask(mask)
			fmt.Printf("Ошибок: %d/%d\n", maxMistakes-mistakes, maxMistakes)

			char, newGuessed := getCharFromUser(guessedChars)
			guessedChars = newGuessed

			isOpened, newMask := openThisChar(char, word, mask)
			mask = newMask
			if !isOpened {
				mistakes--
			}
		}

		printMask(mask)

		if isGameFinished(word, mask, mistakes) {
			if mistakes > 0 {
				fmt.Println("Вы выиграли! Слово:", word)
			}
			if mistakes <= 0 {
				fmt.Println("Вы проиграли! Слово:", word)
			}
		}
	}
}
