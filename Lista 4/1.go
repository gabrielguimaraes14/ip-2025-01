package main

import "fmt"

func main() {
	var numeros [10]int
	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&numeros[i])
	}
	encontrou := false
	fmt.Println("\nNúmeros superiores a 50 e suas posições:")
	for i, num := range numeros {
		if num > 50 {
			fmt.Printf("Número: %d, Posição: %d\n", num, i)
			encontrou = true
		}
	}
	if !encontrou {
		fmt.Println("\nNenhum número superior a 50 foi encontrado.")
	}
}
