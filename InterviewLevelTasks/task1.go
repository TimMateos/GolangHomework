//Условие задачи
//Дан массив целых чисел nums и целое число k. Нужно написать функцию,
//которая вынимает из массива nums k наиболее часто встречающихся элементов.

//Пример
//# ввод
//nums = [1,1,1,2,2,3]
//k = 2
//# вывод (в любом порядке)
//[1, 2]

package main 

import (
    "fmt"
    "sort"
  )

func main() {
  arr := []int{1,1,1,2,2,3}
  k := 2
  fmt.Print(topKFrequentElements(arr, k))
}

func topKFrequentElements(arr []int, k int) []int {
  mHash := make(map[int]int)
  for _, val := range arr {
    mHash[val]++
}

  uniqueNums := make([]int, 0, len(mHash))

	for key := range mHash {
		uniqueNums = append(uniqueNums, key)
	}

  sort.Slice(uniqueNums, func(i, j int) bool {
		return mHash[uniqueNums[i]] > mHash[uniqueNums[j]]
	})
  return uniqueNums[:k]
}
