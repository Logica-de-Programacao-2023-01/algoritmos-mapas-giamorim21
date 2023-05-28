package main

import "fmt"

//Escreva uma função que receba uma string e retorne um mapa onde as chaves
//são os caracteres presentes na string e
//os valores são a frequência de cada caractere.

func caracter(s string) map[string]int {
	caracteres := make(map[string]int)

	for _, char := range s {
		caracteres[string(char)]++
	}

	return caracteres

}

func main() {
	s := "abacaxi"
	resultado := caracter(s)

	fmt.Println(resultado)
}
