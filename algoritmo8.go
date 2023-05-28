package main

import "fmt"

//Escreva uma função que receba um mapa onde as chaves são nomes
//de pessoas e os valores são as quantias de dinheiro que cada
//pessoa gastou em uma conta compartilhada. A função deve calcular
//o valor que cada pessoa deve receber ou pagar para igualar as
//despesas.

func equilibrio(contas map[string]float64) map[string]float64 {
	soma := 0.0
	for _, conta := range contas {
		soma += conta
	}
	media := soma / float64(len(contas))

	resultado := make(map[string]float64)
	for pessoa, conta := range contas {
		diferenca := conta - media
		resultado[pessoa] = diferenca
	}
	return resultado
}

func main() {
	contas := map[string]float64{
		"Giovana": 100.0,
		"Tiago":   50.0,
		"Iza":     75.0,
		"kay":     80.0,
	}
	resultado := equilibrio(contas)
	fmt.Println(resultado)
}
