package main

import (
	"fmt"
)

func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}

func joinWords(word1 string, word2 string) string {
	return fmt.Sprintf("1. %s %s", word1, word2)

}
func joinWords2(word1, word2 string) string {
	return fmt.Sprintf("2. %s %s", word1, word2)

}
func joinWords3(word1, word2 string) (int64, string) {
	return 3, fmt.Sprintf("%s %s", word1, word2)

}

func main() {
	fmt.Println("Main")
	fmt.Println(joinWords("Carlos", "Estrada"))
	fmt.Println(joinWords2("Carlos", "Estrada"))
	number, name := joinWords3("Hola", "Carlos")
	fmt.Println(number, name)
}
