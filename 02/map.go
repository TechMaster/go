package main

import "fmt"

func demoEnglishDictionary() {
	//dict := make(map[string]string)
	dict := map[string]string{}
	dict["hello"] = "xin chào"
	dict["bye"] = "tạm biệt"
	dict["what to eat?"] = "ăn gì?"
	dict["kiss"] = "hôn"
	for key, value := range dict {
		fmt.Println(key, " : ", value)
	}
}

func demoMapInterface() {
	config := map[string]interface{}{
		"host": "localhost",
		"port": 8080,
		"user": "rock",
		"pass": "@123k9",
		"tls":  true,
	}
	for key, value := range config {
		fmt.Println(key, " : ", value)
	}

}
