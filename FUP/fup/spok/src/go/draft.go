package main

import "fmt"

func seEPalindromo(id int) int {
    var numInvertido int

    if id == (((id % 1000 - id % 100) / 100) + (id % 100 - id % 10) + ((id % 10) * 100)) {

        return 1

    }

    return numInvertido
}

func main() {
    var id int

    fmt.Scan(&id)

    fmt.Println(seEPalindromo(id))
}
