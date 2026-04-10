package main

import "fmt"

func verificarAcesso(wifi, login, admin bool) string {

    if !wifi {return "you must connect to wifi"
    }else if !login {return "you need to login first"
    }else if !admin {return "you must to login as admin"}

    return "done"
}

func main() {
    var wifi, login, admin bool

    fmt.Scan(&wifi, &login, &admin)

    fmt.Println(verificarAcesso(wifi, login, admin))
}
