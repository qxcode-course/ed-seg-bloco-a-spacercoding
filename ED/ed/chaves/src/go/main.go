package main

import (
	"fmt"
)

func main() {

	times := NewQueue[rune]()
	ganhadores := NewQueue[rune]()

	for i := range 16 {

		times.Enqueue(rune(65 + i))

	}

	for {
		var x, y int

		fmt.Scan(&x, &y)

		if x > y {

			ganhadores.Enqueue(times.Dequeue())
			times.Dequeue()

		} else {

			times.Dequeue()
			ganhadores.Enqueue(times.Dequeue())

		}

		//fmt.Println(times.String(), ganhadores.String())
		if times.items.Len() == 0 && ganhadores.items.Len() == 1 {break}
		if times.items.Len() == 0 {

			times.items = ganhadores.items
			ganhadores = NewQueue[rune]()

		}

	}

	fmt.Println(string(ganhadores.Dequeue()))

}
