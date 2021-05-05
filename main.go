package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var grid = [][]string{
		{"Y", "C", "Y", "G", "W", "R", "P", "K", "H", "O", "A", "B", "U", "V", "H"},
		{"S", "C", "I", "R", "F", "Z", "B", "M", "C", "P", "M", "Y", "C", "F", "P"},
		{"U", "A", "F", "R", "X", "T", "W", "L", "O", "T", "A", "S", "M", "X", "C"},
		{"E", "J", "R", "A", "G", "S", "A", "V", "H", "G", "L", "R", "X", "G", "F"},
		{"K", "X", "Z", "T", "A", "P", "C", "V", "J", "Q", "M", "J", "Y", "M", "G"},
		{"G", "C", "X", "Q", "E", "W", "S", "I", "A", "L", "A", "E", "O", "I", "V"},
		{"I", "F", "Y", "F", "X", "V", "A", "L", "P", "A", "L", "H", "E", "T", "A"},
		{"L", "E", "K", "O", "U", "U", "T", "I", "G", "U", "A", "N", "C", "O", "I"},
		{"V", "H", "I", "H", "Z", "U", "A", "I", "F", "R", "D", "B", "A", "L", "U"},
		{"A", "R", "Z", "H", "X", "C", "L", "C", "O", "G", "E", "E", "X", "V", "R"},
		{"U", "N", "B", "S", "T", "M", "U", "S", "I", "C", "A", "T", "L", "A", "A"},
		{"W", "R", "A", "U", "J", "A", "B", "I", "S", "S", "N", "O", "R", "I", "S"},
		{"C", "M", "P", "L", "E", "N", "P", "A", "L", "C", "O", "A", "H", "B", "E"},
		{"T", "M", "F", "O", "T", "Z", "M", "P", "T", "R", "E", "S", "J", "R", "L"},
		{"F", "S", "I", "K", "U", "F", "P", "E", "Q", "T", "A", "M", "L", "O", "J"},
	}

	for _, row := range grid {
		for _, grid_element := range row {
			fmt.Printf("%s ", grid_element)
		}
		fmt.Print("\n")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	raw_input, _ := reader.ReadString('\n')

	var input []string

	for _, letter := range raw_input {
		if string(letter) != "\n" {
			input = append(input, string(letter))
			fmt.Println("* ", string(letter))
		}
	}

	fmt.Println(input)
}
