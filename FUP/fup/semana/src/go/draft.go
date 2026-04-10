package main

import "fmt"

func calcHoraTrabalho(dia, hora int) string {

    if (dia == 1  || (hora < 8 || (hora > 11 && hora < 14) || hora > 17)) {return "NAO"
    }else if dia == 7 && (hora < 8 || hora > 11) {return "NAO"}

    return "SIM"
}

func main() {
    var dia, hora int

    fmt.Scan(&dia, &hora)

    fmt.Println(calcHoraTrabalho(dia, hora))
}
