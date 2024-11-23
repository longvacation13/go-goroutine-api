package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const TASK_COUNT = 10000 // 작업 개수

func main() {
	// HTTP 핸들러 등록
	http.HandleFunc("/goroutine-test", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // 시작 시간 기록

		var wg sync.WaitGroup

		// TASK_COUNT 만큼 고루틴 생성
		for i := 0; i < TASK_COUNT; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				performTask(id)
			}(i)
		}

		wg.Wait() // 모든 작업 완료 대기

		duration := time.Since(start) // 종료 시간 기록
		response := fmt.Sprintf("Goroutines completed in: %v\n", duration)
		fmt.Fprintln(w, response)
	})

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil) // HTTP 서버 시작
}

// 고루틴이 실행할 작업
func performTask(id int) {
	time.Sleep(10 * time.Millisecond) // 10ms 대기
}
