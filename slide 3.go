package main

import "fmt"

//Escreva uma função que receba um slice de inteiros e retorne
//um novo slice contendo
//apenas os valores únicos, ou seja, sem duplicatas. Utilize um
//mapa para auxiliar na
//remoção das duplicatas.

func slideNovo(slice []int) []int {
	Map := make(map[int]bool)
	Slice := make([]int, 0)

	for _, n := range slice {
		if !Map[n] {
			Map[n] = true
			Slice = append(Slice, n)
		}
	}
	return Slice
}

func main() {
	n := []int{1, 2, 3, 3, 4, 5, 5, 5, 6, 6, 7, 7, 7, 7, 8, 8, 9}
	
	num := slideNovo(n)
	fmt.Println(num)
}
