package main

import "fmt"

func main() {
    var vetSizeStr int

    fmt.Scan(&vetSizeStr)

    var vetString []string = make([]string, vetSizeStr)

    for i := range vetSizeStr {fmt.Scan(&vetString[i])}

    var vetSizeInt int

    fmt.Scan(&vetSizeInt)

    var vetBusca []string = make([]string, vetSizeInt)

    for i := range vetSizeInt {fmt.Scan(&vetBusca[i])}

    fmt.Println("Hello, World!")

}