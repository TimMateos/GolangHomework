package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// time.Sleep использовать нельзя. это будет не валидным ответом на собеседовании
// Есть функция unpredictableFunc, работающая неопределенно долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).

func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

var ch = make(chan int64) //Создаю канал для записи результата выполнения unpredictableFunc

// Нужно изменить функцию обертку predictableFunc, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.

func predictableFunc() (int64, error) { // добавил возвращение ошибки для контекста
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // создаю контекст для использования таймаута
	defer cancel()

	select { // конструкция select-case для определения, что случилось первым: кончилось время или посчиталась функция
	case <-ctx.Done():
		return 0, ctx.Err() //истечение времени и сообщение об этом
	case x := <-ch:
		return x, fmt.Errorf("Вычисление произведено успешно") // вывод вычисления
	}
}

func main() {
	fmt.Println("started")

	go func() {
		ch <- unpredictableFunc()
	}() // записываю результат функции в канал в отдельной горутине, так как процесс долгий

	fmt.Println(predictableFunc())

	fmt.Println("finished")
}
