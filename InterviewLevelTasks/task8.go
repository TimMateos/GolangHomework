//Написать код, который будет выводить
//коды ответов на HTTP-запросы по двум URL
//главная страница Google и главная страница WB.
//Запросы должны осуществляться параллельно.

//Ответ:

package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	urls := []string{
		"https://www.google.com",
		"https://www.wildberries.ru",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)  // Создаем основной контекст с таймаутом 3 секунды
  defer cancel()

  var wg sync.WaitGroup

  for _, url := range urls {
    wg.Add(1)
    go func(u string) {
      defer wg.Done()
      
      req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
        if err != nil {
				fmt.Print("Ошибка...", err)
				return
			}

      resp, err := http.DefaultClient.Do(req)
         if err != nil {
				fmt.Print("Ошибка...", err)
				return
      }

      defer resp.Body.Close()
      fmt.Printf("URL: %-25s | Status: %s\n", u, resp.Status)

  }(url)

  wg.Wait()
}
}

//если много запросов, то как внедрить воркер пул?
Чтоб работало с воркер пулом надо: 
Создать канал jobs := make(chan string)
Запустить N воркеров, которые постоянно читают из jobs
Затолкать все URL в канал jobs и закрыть его (close(jobs))
Воркеры разберут задачи, и когда канал опустеет — они завершатся.
