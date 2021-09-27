package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	//demoGoRoutineSimple()
	//print_numbers_alphabets()
	//convert_batch_videos_naive()
	//convert_batch_videos()
	convert_batch_videos_ch()
}
