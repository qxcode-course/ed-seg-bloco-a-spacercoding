package main
import "fmt"
func main() {
    var inicio, fim int

    fmt.Scan(&inicio, &fim)

    for i := inicio; i <= fim; i++ {

        if i % 15 == 0 {

            fmt.Println("zigzag")

        } else if i % 3 == 0 {

            fmt.Println("zig")

        } else if i % 5 == 0 {
        
            fmt.Println("zag")

        } else {
        
            fmt.Println(i)
        
        }

    }
}