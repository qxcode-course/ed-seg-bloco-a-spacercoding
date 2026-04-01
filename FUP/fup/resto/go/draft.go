package main
import "fmt"
func main() {
    var num1, num2 int

    fmt.Scan(&num1, &num2)

    fmt.Printf("%d %d\n", num1 / num2, num1 % num2)

}
