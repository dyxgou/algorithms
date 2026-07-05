package main

import (
	"sync"
	"testing"
)

var counter int

func increment() {
	for range 100 {
		counter++
	}
}

func TestSeventh(t *testing.T) {
	res := 100 * 100
	var wg sync.WaitGroup

	wg.Add(100)

	for range 100 {
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()

	t.Logf("counter expected=%d. got=%d", res, counter)
}
