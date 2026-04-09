package main
import "fmt"
func main() {
    var n1, n2, div float64
    
    fmt.Scan(&n1, &n2)
    
    div = n1 / n2
    resto := int(n1) % int(n2)


    fmt.Printf("%d\n%d\n%.2f\n", int(div), resto, div)
}
