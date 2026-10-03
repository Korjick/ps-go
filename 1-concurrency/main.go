package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func producer() <-chan int {
	dataChan := make(chan int)
	go func() {
		defer close(dataChan)

		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			val := rand.Intn(100)
			wg.Add(1)
			go mapper(val, dataChan, &wg)
		}
		wg.Wait()
	}()
	return dataChan
}

func mapper(value int, dataChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	dataChan <- value * value
}

func main() {
	for value := range producer() {
		fmt.Printf("%d ", value)
	}
}
