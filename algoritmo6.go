package main

import "fmt"

//Escreva uma função que receba uma lista de mapas, onde cada mapa contém a
//contagem de palavras de um texto, e retorne um único mapa contendo
//a soma de todas as contagens.

func contar(a []map[string]int) map[string]int {
	soma := 0
	for _, list := range a {
		for _, num := range list {
			soma += num
		}
	}
	mapa := map[string]int{
		"soma": soma,
	}
	return mapa
}

func main() {
	a := []map[string]int{
		{
			"Maze Runner": 3,
		},
		{
			"Harry Potter": 8,
		},
		{
			"Star Wars": 10,
		},
	}
	fmt.Println(contar(a))
}
