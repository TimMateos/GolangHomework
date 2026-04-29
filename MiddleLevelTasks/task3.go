package main

import (
	"fmt"
	"sync"
	"time"
)

const numRequests = 10000

var count int

var m sync.Mutex

func networkRequest() {
	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса.
	m.Lock()                     // мьютекс, чтобы защитить переменную от конкурентной записи
	count++
	m.Unlock()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(numRequests)                // создание waitgroup для корректного завершения всех горутин
	for i := 0; i < numRequests; i++ { // цикл на 10000 итераций инкремента пускает 10 тыс горутин
		go func() {
			defer wg.Done()
			networkRequest() // горутины запускаются и сразу засыпают
		}()
	}

	wg.Wait()          // ждем все горутины
	fmt.Println(count) //печатаем счетчик
}

/* код отработает достаточно быстро за счет параллельной работы горутин и быстрой отработки мьютекса.
Использование атомиков в данном случае было бы быстрее, так как совершается простая операция инкремента,
но при работе с реальными сетевыми запросами лучше конечно использовать мьютекс
*/
