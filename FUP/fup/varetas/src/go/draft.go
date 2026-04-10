package main

import "fmt"

func calcTriangulo(vareta1, vareta2, vareta3 int) string {

    if (vareta1 >= vareta2 + vareta3) || (vareta2 >= vareta1 + vareta3) || (vareta3 >= vareta2 + vareta1) {return "False"}

    return "True"
}

func main() {
    var vareta1, vareta2, vareta3 int 

    fmt.Scan(&vareta1, &vareta2, &vareta3)

    fmt.Println(calcTriangulo(vareta1, vareta2, vareta3))
}
