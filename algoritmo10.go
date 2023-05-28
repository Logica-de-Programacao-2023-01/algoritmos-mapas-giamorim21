package main

import "fmt"

//Escreva uma função que receba um slice de inteiros e retorne
//um mapa onde as chaves são os pares de números encontrados
//no slice e os valores são as frequências de cada par.

func par(slice []int) map[[2]int]int {
	frequencia := make(map[[2]int]int)

	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			pares := [2]int{slice[i], slice[j]}
			frequencia[pares]++
		}
	}
	return frequencia
}

func main() {
	num := []int{1, 2, 3, 2, 4, 1, 3, 4, 2, 3}
	frequencia := par(num)

	for pares, freq := range frequencia {
		fmt.Printf("%v: %d\n", pares, freq)
	}
}
