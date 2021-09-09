package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x     = 0
	mutex sync.Mutex
	w     sync.WaitGroup
)

//--------------
func inc() {
	x = x + 1
}

func demoIncRace() {
	x = 0
	for i := 0; i < 1000; i++ {
		go inc()
	}
	fmt.Println("final value of x", x)
}

//--------------- Increment with Mutex Lock
func incSync() {
	mutex.Lock()
	x = x + 1
	mutex.Unlock()
}
func demoIncSync() {
	x = 0
	for i := 0; i < 1000; i++ {
		go incSync()
	}
	fmt.Println("final value of x", x)
}

//----------Increment with Wait Group
func incSync2(wg *sync.WaitGroup) {
	//wg.Add(1) add wait group ở đây không đúng
	mutex.Lock()
	x = x + 1
	mutex.Unlock()
	wg.Done()
}

func demoIncSync2() {
	x = 0
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incSync2(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

func DemoRaceCondition() {
	fmt.Println("increment without mutex lock and wait group")
	demoIncRace()

	time.Sleep(8000)
	fmt.Println("increment with mutex lock but no wait group")
	demoIncSync()

	fmt.Println("combine mutex lock and wait group")
	time.Sleep(8000)
	demoIncSync2()
}
