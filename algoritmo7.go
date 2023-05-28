package main

import (
	"fmt"
	"strings"
)

//Escreva uma função que receba uma string contendo uma frase e
//retorne um mapa onde as chaves são as palavras encontradas na
//frase e os valores são mapas contendo a contagem de cada letra
//em cada palavra.

func contagem(frase string) map[string]map[string]int {
	resultado := make(map[string]map[string]int)

	palavras := strings.Fields(frase) // divide a frase em palavras

	for _, palavra := range palavras {
		counts := make(map[string]int)
		for _, letra := range palavra {
			letras := string(letra)
			counts[letras]++
		}
		resultado[palavra] = counts
	}
	return resultado
}

func main() {
	frase := "Hoje vou comer banana"
	resultado := contagem(frase)
	fmt.Println(resultado)
}
