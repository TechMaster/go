# Bài tập buổi 06

> Trong buổi 6 chúng ta đã làm quen với Go Routine. Go Routine để chạy các hàm độc lập với nhau (hàm A không cần phải chờ hàm B xong mới thực hiện). Có thể khai báo hàm trực tiếp bằng lệnh `go func()` hoặc gọi `go tên_hàm_cụ_thể()`. Sau đó chúng ta tìm hiểu về sync package dùng để đồng bộ giữa các routine. sync package có những thành phần `once` để chạy hàm một lần, cho singleton pattern, `mutext` dùng để hạn chế mỗi thời điểm chỉ có 1 hàm được đọc hoặc ghi vào dữ liệu, `waitgroup` để đợi một nhóm các hàm hoàn thành.

> Channel dùng để trao đổi thông tin, trả về kết quả hoặc lỗi giữa các go routine. Channel có 2 loại: một chiều (unidirection) và hai chiều (bidirectional). Có thể tạo channel cho một kiểu dữ liệu bất kỳ. Rồi lại tạo buffer cho channel để chứa tạm message gửi vào.

Giờ là phần bài tập tôi thiết kế dành riêng cho các bạn DevOps OneMount nhé.
Chú ý bài tập không phải mượn sức các bạn để xử lý bài toán tôi đang gặp phải. Mỗi ngày tôi chăm chỉ code 12 tiếng xử lý được những vấn đề tôi cần giải quyết. Đây là bài toán thường gặp file khi làm việc với các hệ thống file khác nhau: docx, xlsx, pdf, png, csv.... Nếu gặp khó khăn hay có những câu hỏi gì, các bạn cứ thoải mái bàn luận, trao đổi. Hãy coi tôi như một lập trình viên trong team.

## Tác vụ 1 (5 điểm): Convert tất cả các file video mp4 lưu ở thư mục A ra định dạng Apple HLS ở thư mục B

Việc cần làm:
1. Tải về 2 file binary `ffmpeg` và `ffprobe` ở trang web [https://ffmpeg.org/download.html](https://ffmpeg.org/download.html) để tải về để trong thư mục `/usr/local/bin` là nó thực thi được.
2. Ứng dụng cần đọc vào đường dẫn thư mục chứa mp4 và thư mực chứa kết quả băm từ command line hoặc trong file cấu hình TXT, JSON, YAML... Khuyến cáo nên dùng thư viện [https://github.com/spf13/viper](https://github.com/spf13/viper)
3. Quét tất cả file mp4 để thực hiện. Viết code làm sao để có thể xử lý được tất cả các file mp4 trong khoảng thời gian ngắn nhất có thể
4. Xử lý lỗi nếu gặp file mp4 không đúng định dạng, do người dùng upload lên linh tinh chả hạn
5. File mp4 gốc tên rất tuỳ tiện do đó cần dùng thư viện [https://github.com/matoous/go-nanoid](https://github.com/matoous/go-nanoid) để sinh mã ngẫu nhiên dài đúng 8 ký tự, chú ý chỉ dùng bộ ký tự `A-Za-z0-9` tránh dùng ký tự `-_` để phần mềm đọc file m3u8 hoạt động đúng.
6. Sau khi băm xong hãy dùng phần mềm [https://www.videolan.org/](https://www.videolan.org/) để xem thử kết quả

Toàn bộ code băm tôi đã viết sẵn ở đây rồi.

```go
type VideoInfo struct {
	Duration   int    // Thời lượng in seconds
	Resolution string // Độ phân giải
}

// Truyền vào đường dẫn file video trả về thời lượng và độ phân giải
func GetVideoInfo(filePath string) (videoInfo VideoInfo, err error) {
	var result []byte
	durationArgs := []string{
		"-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1",
		filePath,
	}
	result, err = exec.Command("ffprobe", durationArgs...).Output()
	if err != nil {
		return videoInfo, eris.NewFromMsg(err, "Lỗi chạy ffprobe").InternalServerError()
	}
	durationStr := string(result)
	indexOfDot := strings.Index(durationStr, ".") //Chỉ lấy phần nguyên, bỏ phần thập phân
	if indexOfDot != -1 {
		durationStr = durationStr[:indexOfDot]
	}

	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		return videoInfo, eris.NewFromMsg(err, "Lỗi đọc thời lượng video").InternalServerError()
	}
	videoInfo.Duration = duration

	// Đọc độ phân giải
	resolutionArgs := []string{
		"-v", "error", "-show_entries", "stream=width,height",
		"-of", "csv=p=0:s=x",
		filePath,
	}
	result, err = exec.Command("ffprobe", resolutionArgs...).Output()
	if err != nil {
		return videoInfo, eris.NewFromMsg(err, "Lỗi đọc độ phân giải video").InternalServerError()
	}

	resolution := strings.ReplaceAll(string(result), "\n", "")
	videoInfo.Resolution = resolution

	return videoInfo, nil
}
/*
Hàm này làm 3 việc: 
1. Tạo folder HLS để chưa kết quả băm
2. Chuẩn bị exec.Command để gọi lệnh ffmpeg
3. Gọi hàm executeHashing trong go routine để băm
*/
func hashVideoToHLS(videoID string) error {
	args, hashedFolder, err := makeHlsDir(videoID)
	if err != nil {
		return err
	}

	cmd := exec.Command("ffmpeg", args...)
	cmd.Dir = hashedFolder
	// Hàm này chạy rất lâu do đó phải cho vào go routine để chạy
	go executeHashing(cmd, uploadMeta)
	return nil
}
/*
Tạo thư mục HLS chờ sẵn để băm
*/
func makeHlsDir(videoID string) (args []string, hashedFolder string, err error) {
	/*Tạo thư mục chứa file m3u8 và các file sau khi băm video
	Lệnh băm ffmpeg sẽ chạy ở thư mục này do đó chỉ cần truyền tên file manifestFile m3u8 là đủ
	*/
	hashedFolder = path.Join(HLSDir, videoID)
	if err := os.MkdirAll(hashedFolder, 0777); err != nil {
		return nil, hashedFolder, eris.NewFromMsg(err, "Lỗi tạo thư mục HLS "+videoID)
	}

	manifiestFile := videoID + ".m3u8"

	videoPath := path.Join(UploadDir, videoID+".mp4")

	//Lấy thông tin độ phân giải
	videoInfo, err := GetVideoInfo(videoPath)
	if err != nil {
		return nil, hashedFolder, err
	}
	resolution := strings.Split(videoInfo.Resolution, "x")
	scale := fmt.Sprintf("scale=w=%s:h=%s:force_original_aspect_ratio=decrease", resolution[0], resolution[1])

	hlsKey, _ := gonanoid.New(10)
	hlsVector, _ := gonanoid.New(10)

	return []string{
		"-i", videoPath,
		"-vf", scale, "-c:v", "h264", "-c:a", "aac", "-ar", "48000", "-profile:v", "main", "-crf", "20",
		"-r", "30", "-g", "60", "-maxrate", "5350k", "-bufsize", "7500k", "-b:a", "192k",
		"-hls_enc", "1", "-hls_enc_key", hlsKey, "-hls_enc_iv", hlsVector,
		"-hls_playlist_type", "vod", manifiestFile,
		"-hide_banner", "-loglevel", "quiet", "-progress", "/dev/stdout",
	}, hashedFolder, nil
}

/*
Hàm này thực sự tiến hành băm
*/
func executeHashing(cmd *exec.Cmd, uploadMeta UploadVideoMeta) {
	err := cmd.Start()
	if err != nil {
    //Phải viết hàm báo lỗi ở đây
		return
	}

	//Đoạn này chạy rất lâu !
	err = cmd.Wait()
	if err != nil {
    //Phải viết hàm báo lỗi ở đây
		return
	}
}
```

## Tác vụ 2 (5 điểm): Quan sát thay đổi trong thư mục chứa mp4 để băm.

Hãy dùng thư viện file watcher tuỳ bạn kiếm trên Internet. Khi có một file mp4 được copy vào hãy tiến hành băm file này.
Bỏ qua trường hợp xoá file mp4 hoặc đổi tên, chỉ tập trung xử lý sự kiện file mp4 mới được copy vào.
Có mấy vấn đề bạn cần suy nghĩ xử lý:
1. Tên file mp4 đầu vào có thể là bất kỳ nhưng thư mục băm out put đầu ra được chuẩn bị bằng mã unique ID sinh bởi [https://github.com/matoous/go-nanoid](https://github.com/matoous/go-nanoid)
2. Trong khi ứng dụng đang băm dở, có một số file mp4 copy vào, vậy làm sao xử lý không bỏ lỡ những file này?
3. Làm sao để biết rằng một file mp4 đã được băm xong. Gợi ý có thể tạo một file info.json trong mỗi thư mục HLS để lưu tên file mp4 gốc để từ đó tránh băm lại hoặc ghi tiến trình băm ra một file để đối chiếu file mp4 gốc và thư mục HLS đã băm.
4. Trong quá trình băm, ứng dụng bị crash, khi khởi động lại, làm sao chỉ băm đúng những file mp4 chưa băm và băm lại file đang xử lý dở dang

