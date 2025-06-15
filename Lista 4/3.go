package main
import "fmt"

func main() {
    vetor := make ([]int, 10)
    fmt.Print("Digite os 10 números do vetor: ")
    for i := 0 ; i < 10 ; i++ {
        fmt.Scan(&vetor[i])
    }
    soma_pares := 0
    quantidade_impares := 0
    fmt.Print("Os números pares digitados são: ")
    for i := 0 ; i < 10 ; i++ {
        if vetor[i] % 2 == 0 {
            fmt.Print(vetor[i], " ")
            soma_pares = soma_pares + vetor[i]
        }
    }
    fmt.Println()
    fmt.Print("A soma dos números pares digitados é ", soma_pares, "\n")
    fmt.Print("Os números ímpares digitados são: ")
    for i := 0 ; i < 10 ; i++ {
        if vetor[i] % 2 != 0 {
            fmt.Print(vetor[i], " ")
            quantidade_impares = quantidade_impares + 1
        }
    }
    fmt.Println()
    fmt.Print("A quantidade de números ímpares digitados é ", quantidade_impares)
}