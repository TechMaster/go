package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Tại sao không trả về error trong go routine?
https://stackoverflow.com/questions/20945069/catching-return-values-from-goroutines
*/
func convert_video_mp4_to_hls(input string, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Println("Convert video " + input)
	wg.Done()
}

func convert_batch_videos() {
	videos := []string{"v1.mp4", "v2.mp4", "v3.mp4", "v4.mp4"}
	wait_group := sync.WaitGroup{}
	for _, video := range videos {
		wait_group.Add(1)
		go convert_video_mp4_to_hls(video, &wait_group)
	}
	wait_group.Wait()
	fmt.Printf("Convert xong %d video\n", len(videos))
}
