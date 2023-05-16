package main

import "fmt"

//Escreva uma função que receba uma string como entrada
//e retorne um mapa onde as
//chaves são os caracteres presentes na string e os
//valores são o número de ocorrências
//de cada caractere.

func contarCaracteres(palavra string) map[string]int {

	caracteres := make(map[string]int)

	for _, caracter := range palavra {
		caracteres[string(caracter)]++
	}
	return caracteres
}

func main() {
	palavra := "banana"
	fmt.Println(contarCaracteres(palavra))

}
