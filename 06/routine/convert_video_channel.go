package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
https://golangbyexample.com/channel-function-argument-go/
*/
type video_convert_result struct {
	input string
	err   error
}

func convert_video_mp4_to_hls_ch(input string, wg *sync.WaitGroup, ch chan<- video_convert_result) {
	defer wg.Done() //defer để chạy hàm này vào lúc cuối trước khi hàm thoát

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	var err error
	if rand.Intn(2) == 1 {
		err = errors.New("Lỗi convert " + input)
	} else {
		err = nil
	}
	ch <- video_convert_result{
		input: input,
		err:   err,
	}
}

func convert_batch_videos_ch() {
	videos := []string{"v1.mp4", "v2.mp4", "v3.mp4", "v4.mp4", "v5.mp4", "v6.mp4", "v7.mp4", "v8.mp4"}
	convert_result_channel := make(chan video_convert_result)
	handle_convert_error(convert_result_channel)

	wait_group := sync.WaitGroup{}
	for _, video := range videos {
		wait_group.Add(1)
		go convert_video_mp4_to_hls_ch(video, &wait_group, convert_result_channel)
	}

	wait_group.Wait()
	fmt.Printf("Convert xong %d video\n", len(videos))
}

/*
Receive channel để xử lý thông báo lỗi
*/
func handle_convert_error(ch <-chan video_convert_result) {
	go func() {
		for {
			convert_result, more := <-ch //khi không còn dữ liệu, more sẽ false
			if convert_result.err != nil {
				fmt.Println(convert_result.err.Error())
			} else {
				fmt.Println("convert successfully " + convert_result.input)
			}

			if !more {
				fmt.Println("No more result")
				return //thoát khỏi go routine
			}
		}
	}()

}
