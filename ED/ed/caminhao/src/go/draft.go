package main
import "fmt"

type Node *struct {

    value struct{int int}
    next *Node
    previous *Node
    head *Node

}

type ListaC *struct {

    head *Node

}

func main() {
    var nLinhas int

    fmt.Scan(&nLinhas)

    var lista ListaC

    for range nLinhas {

        

    }
}