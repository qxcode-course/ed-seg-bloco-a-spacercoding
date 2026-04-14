package main
import "fmt"
func main() {
    var num1, num2 int

    fmt.Scan(&num1, &num2)

    fmt.Print("[ ")
    for i := num1; ; {

        fmt.Printf("%d ", i)
        if num1 > num2 {

            i--

        } else {i++}
        if num1 > num2 && i == num2 {break
        }else if num1 < num2 && i == num2 {break}

    }
    fmt.Println("]")
}
