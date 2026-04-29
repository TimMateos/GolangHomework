package main

import (
	"fmt"
	"sync"
)

func balance() int {
	x := make(map[int]int, 5) // указать размер можно сразу, так как мы знаем, что значений будет 5
	var m sync.Mutex
	var wg sync.WaitGroup // чтоб решить проблему нулей необходимо дождаться выполнения всех горутин

	// call bank
	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()

			b := bank_network_call(i)

			m.Lock() // лок мьютекса ставить надо только на само вычисление, иначе будет сильно увеличено время работы
			x[i] = b
			m.Unlock()
		}()
	}
	wg.Wait() // Ждем все вычисления

	// Как-то считается сумма значений в мапе и возвращается
	return sumOfMap
}
