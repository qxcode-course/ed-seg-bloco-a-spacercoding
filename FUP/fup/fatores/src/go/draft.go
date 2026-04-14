package main
import "fmt"
func main() {
    var num, cont int

    fmt.Scan(&num)

    for i := 2; i <= num; {

        if num % i == 0 {

            cont++
            num = num / i

        } else {

            if cont != 0 {fmt.Println(i, cont)}
            cont = 0
            i++

        }

        if num == 1 {

            fmt.Println(i, cont)
            break

        }

    }
}
