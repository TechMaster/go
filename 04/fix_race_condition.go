package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	x     int
	mutex sync.Mutex
	//Không để waitgroup trong struct này
}

func (c *Counter) inc(wg *sync.WaitGroup) {
	c.mutex.Lock()
	c.x = c.x + 1
	c.mutex.Unlock()
	wg.Done()
}

func DemoCounter() {
	w := sync.WaitGroup{}

	c := Counter{
		x:     0,
		mutex: sync.Mutex{},
	}
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go c.inc(&w)

	}
	w.Wait()
	fmt.Println("final value of x", c.x)
}
