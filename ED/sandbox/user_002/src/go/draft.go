package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type multiSet struct {
	dados      []int
	capacidade int
	tamanho    int
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func NewMultiSet(valor int) *multiSet {

	//novo multiSet com capacidade "valor" e tamanho 0
	return &multiSet{make([]int, valor), valor, 0}

}

func (ms *multiSet) Search(valor int) (bool, int) {
	//indexes usados para a comparação
	inicio := 0
	fim := len(ms.dados)
	meio := (inicio + fim) / 2

	//busca propriamente dita, parando quando não existir mais meio.
	for fim < inicio {

		//condição de parada:
		if valor == ms.dados[meio] {
			return true, meio
		}

		//vendo qual metade vai ser procurada
		if valor > ms.dados[meio] {
			inicio = meio
		} else {
			fim = meio
		}

		//novo meio
		meio = (inicio + fim) / 2
	}

	//se não encontar, quando sai do for:
	return false, meio
}

func (ms *multiSet) expand() {
	//criando vetor novo e expandido
	var dataAux []int = make([]int, ms.capacidade+1)

	//passando as informações pra ele
	for i := range ms.tamanho {

		dataAux[i] = ms.dados[i]

	}

	//transferindo pro multiSet
	ms.dados = dataAux

	//aumentando a informação da capacidade
	ms.capacidade++

}
func (ms *multiSet) insert(valor, index int) {
    if ms.tamanho == 0 {
        ms.dados[0] = valor
        ms.tamanho++
        return
    }

	//abrindo espaço para o novo valor
	for i := ms.tamanho; i > index; i-- {

		ms.dados[i] = ms.dados[i-1]

	}
	
	//inserindo
	ms.dados[index] = valor
	ms.tamanho++

}
func (ms *multiSet) Insert(valor int) {
	if ms.tamanho == 0 {
		ms.insert(valor, 0)
		return
	}
	//caso precise expandir a capacidade
	if ms.capacidade == ms.tamanho {

		ms.expand()

	}

	achou, index := ms.Search(valor)

	//se o número estiver presente
	if achou {
		ms.insert(valor, index)
		return
		//do contrário
	} else {

		for i := 0; i < ms.tamanho; i++ {

			if ms.dados[i] > valor || ms.dados[i] == valor {
				ms.insert(valor, i)
				return
			} else if ms.dados[ms.tamanho-1] < valor {
				ms.insert(valor, ms.tamanho-1)
				return
			}
		}
	}

}

func (ms *multiSet) Clear() {
	var vetVazio []int = make([]int, ms.capacidade)

	ms.dados = vetVazio

}

func (ms *multiSet) String() string {
	var str string = "["

	if ms.tamanho != 0 {
		for i := 0; i < ms.tamanho-1; i++ {
			if ms.dados[i] != 0 {
				str += strconv.Itoa(ms.dados[i]) + ", "
			}
		}
		if ms.dados[ms.tamanho-1] != 0 {
			str += strconv.Itoa(ms.dados[ms.tamanho-1])
		}
	}
	str += "]"

	return str
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			//fmt.Println(ms.dados)
			fmt.Println(ms.String())
		case "erase":
			//value, _ := strconv.Atoi(args[1])
		case "contains":
			// value, _ := strconv.Atoi(args[1])
		case "count":
			// value, _ := strconv.Atoi(args[1])
		case "unique":
			// fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
