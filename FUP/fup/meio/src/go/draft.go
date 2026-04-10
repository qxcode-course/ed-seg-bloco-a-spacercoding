package main

import "fmt"

func doMeio(val1, val2, val3 int) int {

    if (val1 < val2 && val1 > val3) || (val1 > val2 && val1 < val3) {return val1
    }else if (val2 < val1 && val2 > val3) || (val2 > val1 && val2 < val3) {return val2
    }else if (val3 < val1 && val3 > val2) || (val3 > val1 && val3 < val2) {return val3}

    return val1
}

func main() {
    var val1, val2, val3 int

    fmt.Scan(&val1, &val2, &val3)

    fmt.Println(doMeio(val1, val2, val3))
}
