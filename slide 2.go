package main

import "fmt"

//Escreva uma função que receba um mapa onde as chaves são nomes de
//alunos e os
//valores são slices de notas. A função deve calcular a média das
//notas de cada aluno e
//retornar um novo mapa onde as chaves são os nomes dos alunos
//e os valores são as
//médias correspondentes.

func calcularMedia(notas map[string][]float64) map[string]float64 {
	medias := make(map[string]float64)

	for aluno, notas := range notas {
		var total float64 = 0
		
		for _, nota := range notas {
			total += nota
		}
		media := total / float64(len(notas))
		medias[aluno] = media
	}

	return medias
}

func main() {
	notas := map[string][]float64{
		"Giovana": []float64{10.0, 10.0, 10.0},
		"Iza":     []float64{10.0, 10.0, 10.0},
		"Mate":    []float64{10.0, 10.0, 10.0},
	}
	medias := calcularMedia(notas)

	for aluno, media := range medias {
		fmt.Println(aluno, media)
	}
}
