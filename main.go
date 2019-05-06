package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func scanWords(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
}

func mapWords(words []string) map[string][]string {
	wordsMap := make(map[string][]string)

	for i, word := range words {
		key := word
		if i == len(words)-1 {
			fmt.Printf("On last iteration: %d\n", i)
		} else {
			value := words[i+1]
			wordsMap[key] = append(wordsMap[key], value)
		}
	}

	return wordsMap
}

func getKey(m map[string][]string) string {
	for k := range m {
		return k
	}
	return ""
}

func seed(max int, min int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func output(wordsMap map[string][]string, iterations int) string {
	key, followingWord := getKey(wordsMap), ""
	sentence := key + " "

	for i := 0; i <= iterations; i++ {
		if len(wordsMap[key]) > 1 {
			index := seed(len(wordsMap[key]), 0)
			followingWord = wordsMap[key][index]
		} else {
			followingWord = wordsMap[key][0]
		}
		sentence += followingWord + " "
		key = followingWord
	}

	return sentence
}

func main() {
	words, err := scanWords("test.txt")
	if err != nil {
		panic(err)
	}

	wordsMap := mapWords(words)
	//fmt.Printf("\nWords Map: %s", wordsMap)

	randVal := output(wordsMap, 100)
	fmt.Printf("\n\nOutput: %s", randVal)
}
