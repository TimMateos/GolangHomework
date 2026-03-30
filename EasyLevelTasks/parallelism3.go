package main

import (
	"fmt"
	"sync"
	"time"
)

/*
В данном задании происходил data race при записи в мапу и GO останавливал программу
Так как мапы не потокобезопасны, то при попытке горутин записывать что-то в одну и ту мапу одновременно
приводит к ошибке concurrent map writes
*/

func main() {
	x := make(map[int]int, 1)
	mu := sync.Mutex{}
	write := func(key, value int) {
		mu.Lock()
		x[key] = value
		mu.Unlock()
	}
	go func() { write(1, 2) }()
	go func() { write(3, 7) }()
	go func() { write(123, 10) }()
	go func() { write(1, 2) }()
	go func() { write(34, 7) }()
	go func() { write(1432, 10) }()
	go func() { write(1, 2) }()
	go func() { write(100, 7) }()
	go func() { write(34, 10) }()
	go func() { write(1, 2) }()

	time.Sleep(100 * time.Millisecond) //блокируемся на 100 миллисекунд

	fmt.Println("x[1] =", x[1])
}
