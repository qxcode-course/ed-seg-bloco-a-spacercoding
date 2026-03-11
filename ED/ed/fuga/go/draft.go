package main

import "fmt"

func imprimir(resultado string) {

	fmt.Print(resultado)

}

func pegaLadrao(posHeli, posPol, posFug, direcao int) string {

	if direcao == 1 {

		if posHeli > posPol && posFug > posPol {
			return "S\n"
		} else if posHeli < posPol && posFug > posPol {
			return "S\n"
		} else {
			return "N\n"
		}

	} else {

		if posHeli > posPol && posFug < posPol {
			return "S\n"
		} else if posHeli < posPol && posFug < posPol {
			return "S\n"
		} else if posHeli > posPol && posFug > posHeli {
			return "S\n"
		} else {
			return "N\n"
		}

	}

}

func alocacao(posHeli, posPol, posFug, direcao *int) {

	fmt.Scan(posHeli, posPol, posFug, direcao)

}

func main() {

	var posHeli, posPol, posFug, direcao int

	alocacao(&posHeli, &posPol, &posFug, &direcao)

	imprimir(pegaLadrao(posHeli, posPol, posFug, direcao))

}
