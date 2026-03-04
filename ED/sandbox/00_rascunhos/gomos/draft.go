package main

import "fmt"

type posicaoGomo struct {

    x, y int

}

func atualizarPosicaoCobra(cobra []posicaoGomo, direcao string, quantidade int) ([]posicaoGomo) {

    if direcao == "L" {

        for i := range quantidade {
            if i == 0 {continue
            }else {cobra[i] = cobra[i-1]}
        }

        cobra[0].x--

    }else if direcao == "R" {

        for  i := len(cobra)-1; i > 1; i-- {
            cobra[i] = cobra[i-1]
        }
        
        cobra[0].x++

    }else if direcao == "U" {

        for i := range quantidade {
            if i == 0 {continue
            }else {cobra[i] = cobra[i-1]}
        }
        
        cobra[0].y--

    }else if direcao == "D" {

        for i := range quantidade {
            if i == 0 {continue
            }else {cobra[i] = cobra[i-1]}
        }
        
        cobra[0].y++

    }

    return cobra

}

func main() {

    var quantidade int
    var direcao string
    fmt.Scan(&quantidade, &direcao)
    
    var cobra = make([]posicaoGomo, quantidade)
    
    for i := range quantidade {
        
        fmt.Scan(&cobra[i].x, &cobra[i].y)
        
    }
    
    atualizarPosicaoCobra(cobra, direcao, quantidade)
    
    for i := range cobra {

        fmt.Println(cobra[i].x, cobra[i].y)

    }

}
