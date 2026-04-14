package main
import "fmt"
func main() {
    var num1, num2 int

    fmt.Print("[ ")
    fmt.Scan(&num1, &num2)
    for i := num1; i < num2; i++ {

        fmt.Printf("%d ", i)

    }
    fmt.Println("]")
}
