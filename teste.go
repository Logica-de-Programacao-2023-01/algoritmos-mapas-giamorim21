package main

import "fmt"

//func main() {
//	crescente := []int{1, 3, 4, 6, 8, 7, 5, 9, 2}
//	sort.Ints(crescente)
//	fmt.Println(crescente)
//}

func main() {
	exemplo := make(map[string]int)

	exemplo["maçã"] = 10
	exemplo["banana"] = 15
	exemplo["abacaxi"] = 20

	quantidadeMaca := exemplo["maçã"]

	exemplo["maçã"] = 20

	delete(exemplo, "maçã")

	_, ok := exemplo["maçã"]
	if ok {
		fmt.Println("existe no mapa")
	} else {
		fmt.Println("não existe no mapa")
	}

	fmt.Println(exemplo)
	fmt.Println(quantidadeMaca)
}
