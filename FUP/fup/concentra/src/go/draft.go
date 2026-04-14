package main
import "fmt"
func main() {
    var num1, num2, index2 int

    fmt.Scan(&num1, &num2)

    fmt.Print("[ ")
    index2 = num2
    for i := num1; i <= num2; i++ {

        fmt.Printf("%d %d ", i, index2)
        index2--

    }
    fmt.Println("]")
}
