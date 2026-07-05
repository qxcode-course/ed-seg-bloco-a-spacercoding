package main

import (
	"fmt"
)

func result(qT *Queue[rune], qG *Queue[int]) rune {

	for range 15 {

		if qG.items.Back().Value > qG.items.Back().Prev().Value {



		}

	}

}

func main() {

	qT := NewQueue[rune]()
	qG := NewQueue[int]()

	for i := range 8 {

		qT.Enqueue(rune(97 + i))

	}

	for range 15 {
		var x, y int

		fmt.Scan(&x, &y)

		qG.Enqueue(x)
		qG.Enqueue(y)

	}

	fmt.Println(result(qT, qG))

}
