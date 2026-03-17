package main

import (
    "fmt"
    "slices"
    "strconv"
)

func atribuirInt(inteiro *int) {

    fmt.Scan(inteiro)

}

func preencherVetorInt(vetor []int) {

    for i := range vetor {

        vetor[i] = i+1

    }

}

func josephus(termoInicio int, jogo []int) {

    if len(jogo) > 2 {

        toString(jogo)

        for i := range jogo {

            if jogo[i] == termoInicio && termoInicio < len(jogo) - 2 {

                jogo = slices.Delete(jogo, i, i + 1)
                termoInicio += 2

            } else if jogo[i] == termoInicio && termoInicio == len(jogo) - 2 {

                jogo = slices.Delete(jogo, len(jogo)-2, len(jogo)-1)
                termoInicio = jogo[0]

            } else if jogo[i] == termoInicio && termoInicio > len(jogo) - 2 {

                jogo = slices.Delete(jogo, 0, 1)
                termoInicio = jogo[0]

            }

        }

        josephus(termoInicio, jogo)

    } else if len(jogo) == 2{

        toString(jogo)
        if jogo[0] == termoInicio {

            jogo = slices.Delete(jogo, 0, 1)
            toString(jogo)

        } else {

            jogo = slices.Delete(jogo, 0, 0)
            toString(jogo)

        }

    }

}

func toString(vetor []int) {

    var retorno = make([]string, 0)

    for i := range vetor {
        if vetor[i] > 0 {
            retorno = append(retorno, strconv.Itoa(vetor[i]))
        }

    }

    fmt.Printf("[ ")
    for i := range retorno {
        fmt.Printf("%s ", retorno[i])
    }
    fmt.Println("]")

}

func main() {

    var qtdPessoas, termoInicio int

    atribuirInt(&qtdPessoas)
    atribuirInt(&termoInicio)

    var jogo = make([]int, qtdPessoas)

    preencherVetorInt(jogo)

    josephus(termoInicio, jogo)

}
