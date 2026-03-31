package main

import (
	"fmt"
    "math/rand"
)

func compararJogada(jogadaJog, jogadaPc, ptsJog, ptsPc int) {

    if  jogadaJog == 1 && jogadaPc == 3 {
        fmt.Printf("Você jogou PEDRA e o PC TESOURA.\nVocê ganhou!\n")
        ptsJog++
    } else if jogadaJog == 1 && jogadaPc == 2 {
        fmt.Printf("Você jogou PEDRA e o PC PAPEL .\nO PC ganhou!\n")
        ptsPc++
    } else if jogadaJog == 2 && jogadaPc == 1 {
        fmt.Printf("Você jogou PAPEL e o PC PEDRA .\nVocê ganhou!\n")
        ptsJog++
    } else if jogadaJog == 2 && jogadaPc == 3 {
        fmt.Printf("Você jogou PAPEL e o PC TESOURA .\nO PC ganhou!\n")
        ptsPc++
    } else if jogadaJog == 3 && jogadaPc == 1 {
        fmt.Printf("Você jogou TESOURA e o PC PEDRA .\nO PC ganhou!\n")
        ptsPc++
    } else if jogadaJog == 3 && jogadaPc == 2 {
        fmt.Printf("Você jogou TESOURA e o PC PAPEL.\nVocê ganhou!\n")
        ptsJog++
    } else {
        fmt.Printf("Você jogou IGUAL ao PC  .\nNinguém ganhou!\n")
    }

}

func main() {
    var jogadaJog, jogadaPc, ptsJog, ptsPc, numRound int
    ptsJog, ptsPc, numRound = 0, 0, 1

    
    
    for numRound; numRound < 6; numRound++{
    fmt.Printf("# JOKENPÔ #\nVocê: %d | PC: %d\nRound: %d/5\n\n1 - Pedra\n2 - Papel\n3 - Tesoura\n", ptsJog, ptsPc, numRound)
    fmt.Scan(&jogadaJog)
    jogadaPc = rand.Intn(4)
    compararJogada(jogadaJog, jogadaPc, ptsJog, ptsPc)
    if numRound == 5 {

        fmt.Println("PLACAR FINAL:\nVocê: %d | PC: %d\n\nJOGAR NOVAMENTE?\n1 - Sim\n0 - Sair\n", ptsJog, ptsPc)
        fmt.Scan(&jogadaJog)
        if jogadaJog == 1 {
            numRound = 0
        } else {
            break
        }

    }
    }

}
