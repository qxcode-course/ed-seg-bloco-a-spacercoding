package main

import "fmt"

func main() {
	var preco, qtdParcelas float64

	fmt.Scan(&preco, &qtdParcelas)

    parcela := preco / qtdParcelas
    custoParcela :=parcela * ((qtdParcelas - 1) * 0.05)

	if qtdParcelas == 1 {
		fmt.Printf("%.2f\n%.2f\n", preco, preco)
	} else {
		fmt.Printf("%.2f\n%.2f\n", parcela + custoParcela, (parcela +custoParcela) * qtdParcelas )
	}
}
