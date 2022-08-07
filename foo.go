package main

import (
	"C"
	"sync"
)

//export Add
func Add(a, b int) int {
	wg := sync.WaitGroup{}
	wg.Add(1)
	result := 0
	// you can still use goroutines
	go func() {
		result = a + b
		wg.Done()
	}()
	wg.Wait()
	return result
}

func main() {}
