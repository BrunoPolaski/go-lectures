package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("concurrency")
		wg.Done()
	}()
	wg.Wait()
}