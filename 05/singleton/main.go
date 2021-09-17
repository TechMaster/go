package main

import "singleton-pattern/config"

func main() {
	// ins1 := singleton.GetInstanceFunc()
	// ins2 := singleton.GetInstanceFunc()

	// fmt.Printf("%p\n", ins1)
	// fmt.Printf("%p\n", ins2)

	//singleton.DemoGetInstanceFunc()
	//singleton.DemoGetInstanceCheck()
	//singleton.DemoGetInstanceMutex()
	//time.Sleep(time.Second * 5)

	config.DemoConfig()
	// database.DemoDatabase()
	//post.DemoPost()
}