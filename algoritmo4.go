package main

import (
	"fmt"
	"sort"
	"strings"
)

//Escreva uma função que receba um slice de palavras e retorne
//um mapa onde as chaves são palavras que são anagramas entre si
//e os valores são os grupos de palavras anagramas.

func anagrama(palavras []string) map[string][]string {
	grupos := make(map[string][]string) // cria um map vazio para armazenar os grupos de palavras

	for _, palavra := range palavras { // percorre cada palavra da slice palavras []string
		sortPalavra := sortString(palavra)
		grupos[sortPalavra] = append(grupos[sortPalavra], palavra)
	}
	return grupos
}
func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	palavras := []string{"listen", "silent", "enlist", "dealing", "leading", "alignment"}
	grupos := anagrama(palavras)

	for _, g := range grupos {
		fmt.Println(g)
	}
}
