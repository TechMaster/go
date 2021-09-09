package main

import (
	"fmt"
	"time"
)

type Config struct {
	Host string
	Port int
}

var config = Config{ //kiểu value
	Host: "localhost",
	Port: 8080,
}

var pConfig = new(Config) //Kiểu con trỏ

func init() {
	pConfig.Host = "127.0.0.1"
	pConfig.Port = 9000
}
func useConfigNormalStyle() {

	fmt.Println("Configure #1")
	fmt.Println(config.Host, config.Port)

	fmt.Println("Configure #2")
	fmt.Println(pConfig.Host, pConfig.Port)

}

/*
Hãy thử chạy hàm useConfig
*/
func useConfigGoStyle() {
	go func() {
		fmt.Println("Configure #1")
		fmt.Println(config.Host, config.Port)
	}()

	go func() {
		fmt.Println("Configure #2")
		fmt.Println(pConfig.Host, pConfig.Port)
	}()
	time.Sleep(1 * time.Second)
}
