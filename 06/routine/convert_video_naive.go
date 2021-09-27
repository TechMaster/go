package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Tại sao không trả về error trong go routine?
https://stackoverflow.com/questions/20945069/catching-return-values-from-goroutines
*/
func convert_video_mp4_to_hls_naive(input string) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Println("Convert video " + input)
}

func convert_batch_videos_naive() {
	videos := []string{"v1.mp4", "v2.mp4", "v3.mp4", "v4.mp4"}
	for _, video := range videos {
		go convert_video_mp4_to_hls_naive(video)
	}
	fmt.Printf("Convert xong %d video\n", len(videos))
}
