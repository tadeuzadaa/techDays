package main

import "fmt"

func main() {
    var nome string
    var idade int

    fmt.Print("Digite seu nome: ")
    fmt.Scanln(&nome)

    fmt.Print("Digite sua idade: ")
    fmt.Scanln(&idade)

    fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
}

