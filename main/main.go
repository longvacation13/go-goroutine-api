package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

const TASK_COUNT = 10000 // 작업 개수

// 메모리 사용 체크 함수
func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v KiB, TotalAlloc = %v KiB, Sys = %v KiB, NumGC = %v\n",
		m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}

// HTTP 핸들러 등록 및 고루틴 성능 테스트
func main() {
	http.HandleFunc("/goroutine-test", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // 시작 시간 기록

		fmt.Println("\nBefore starting goroutines:")
		printMemUsage() // 작업 전 메모리 사용량

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

		fmt.Println("\nAfter finishing goroutines:")
		printMemUsage() // 작업 후 메모리 사용량

		// 강제로 GC 호출 후 메모리 확인
		runtime.GC()
		fmt.Println("\nAfter garbage collection:")
		printMemUsage()
	})

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil) // HTTP 서버 시작
}

// 고루틴이 실행할 작업
func performTask(id int) {
	time.Sleep(10 * time.Millisecond) // 10ms 대기
}
