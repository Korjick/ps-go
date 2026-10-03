package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func producer(dataChan chan<- int) {
	data := make([]int, 10)
	var wg sync.WaitGroup
	for i := range data {
		wg.Add(1)
		value := rand.Intn(100)
		data[i] = value
		go mapper(value, dataChan, &wg)
	}
	wg.Wait()
	close(dataChan)
}

func mapper(value int, dataChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	dataChan <- value * value
}

func main() {
	dataChan := make(chan int)
	go producer(dataChan)
	for value := range dataChan {
		fmt.Printf("%d ", value)
	}
}
