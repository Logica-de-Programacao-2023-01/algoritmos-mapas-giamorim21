package main

import (
	"fmt"
	"strings"
)

//Escreva uma função que conte a ocorrência de cada palavra em uma
//string e retorne um mapa onde as chaves são as palavras
//encontradas e os valores são o número de ocorrências de
//cada palavra.

func contarPalavras(texto string) map[string]int {
	palavras := strings.Fields(texto)
	ocorrencias := make(map[string]int)

	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}
	return ocorrencias
}

func main() {
	p := "Hoje foi para o Ceub"
	fmt.Println(contarPalavras(p))
}
