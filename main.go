package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	linha  int
	coluna int
}

func verifica(input []string, grid [][]string, linha int, coluna int, dir_linha int, dir_coluna int) bool {
	for i, input_letter := range input {
		if coluna+i*dir_coluna < 0 ||
			coluna+i*dir_coluna > 14 ||
			linha+i*dir_linha < 0 ||
			linha+i*dir_linha > 14 ||
			input_letter != grid[linha+i*dir_linha][coluna+i*dir_coluna] {
			return false
		}
	}
	return true
}

func print_grid(grid [][]string) {
	for _, row := range grid {
		for _, grid_element := range row {
			fmt.Printf("%s ", grid_element)
		}
		fmt.Print("\n")
	}
}

func revela(answer_grid [][]string, word []string, i int, j int, dir_linha int, dir_coluna int) {
	for contador, letter := range word {
		answer_grid[i+contador*dir_linha][j+contador*dir_coluna] = letter
	}
}

func inicializa_answer_grid(answer_grid [][]string) [][]string {
	for i := 0; i < 15; i++ {
		answer_grid = append(answer_grid, []string{})
		for j := 0; j < 15; j++ {
			answer_grid[i] = append(answer_grid[i], "*")
		}
	}
	return answer_grid
}

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

	print_grid(grid)

	reader := bufio.NewReader(os.Stdin)

	for {
		var input []string
		fmt.Print("Enter text: ")
		raw_input, _ := reader.ReadString('\n')
		for _, letter := range raw_input {
			if string(letter) != "\n" {
				input = append(input, string(letter))
			}
		}

		var answer_grid = [][]string{}
		answer_grid = inicializa_answer_grid(answer_grid)

		solucoes := []Coord{}

		for i, row := range grid {
			for j, grid_element := range row {
				if grid_element == input[0] {
					for dir_linha := -1; dir_linha < 2; dir_linha++ {
						for dir_coluna := -1; dir_coluna < 2; dir_coluna++ {
							if verifica(input, grid, i, j, dir_linha, dir_coluna) {
								solucoes = append(solucoes, Coord{i, j})
								revela(answer_grid, input, i, j, dir_linha, dir_coluna)
							}
						}
					}
				}
			}
		}

		print_grid(answer_grid)

		for _, solucao := range solucoes {
			fmt.Println("Encontrado: ", solucao)
		}
	}

}
