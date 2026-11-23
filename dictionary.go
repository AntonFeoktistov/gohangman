package main

import (
	"bufio"
	"math/rand/v2"
	"os"
	"strings"
)

func readWordsFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			words = append(words, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return words, nil
}

func getRandomWord(words []string) string {
	len := len(words)
	randomIndex := rand.IntN(len)
	return words[randomIndex]
}
