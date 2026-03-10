package main

import (
    "fmt"
    "slices"
)

func josephus(circuloPessoas []int, numInicio int) ([]int) {

    
    if len(circuloPessoas) > 1 {
        fmt.Println(circuloPessoas)

        for i := 0; i < len(circuloPessoas)-1; i++ {

            if circuloPessoas[i] == numInicio && i < len(circuloPessoas) - 1 {

                circuloPessoas = slices.Delete(circuloPessoas, i+1, i+2)
                josephus(circuloPessoas, circuloPessoas[i])

            } else if i == numInicio && i == len(circuloPessoas) - 1{

                circuloPessoas = slices.Delete(circuloPessoas, 1, 2)
                josephus(circuloPessoas, circuloPessoas[0])

            }

        }

    } 

        return circuloPessoas

}

func main() {

    var numPessoas, numInicio int
    
    fmt.Scan(&numPessoas, &numInicio)

    var circuloPessoas = make([]int, numPessoas)

    for i := 0; i < numPessoas; i++{

        circuloPessoas[i] = i + 1

    }

    josephus(circuloPessoas, numInicio)

}
