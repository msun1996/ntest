package ntest

import (
	"fmt"
	"sync"
	"testing"
)

/*
1. go 编写一个并发池
2. go 两个协程打印出a1b2c3d4.....
*/

func worker(workNum int, works <-chan int, results chan<- int) {
	for work := range works {
		result := work * 2
		fmt.Printf("workNum: %d, work: %d\n", workNum, work)
		results <- result
	}
}

func TestGo(t *testing.T) {
	jobNum := 1000
	// 工作线程数
	workNum := 10

	works := make(chan int, jobNum)
	results := make(chan int, jobNum)

	wg := sync.WaitGroup{}

	for i := 0; i < workNum; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(i, works, results)
		}(i)
	}

	for i := 0; i < jobNum; i++ {
		works <- i
	}
	// 关闭works确保执行完成，各go正常退出
	close(works)

	go func() {
		// go 不全部退出，wg 会永远阻塞。
		wg.Wait()
		close(results)
	}()

	// 阻塞函数，直到 close results
	for result := range results {
		fmt.Printf("results: %d\n", result)
	}
}

func TestGo2(t *testing.T) {
	letter := make(chan bool)
	num := make(chan bool)
	done := make(chan bool)

	go func() {
		for i := 'a'; i <= 'z'; i++ {
			<-letter
			fmt.Print(string(i))
			num <- true
		}
	}()

	go func() {
		for i := 1; i <= 26; i++ {
			<-num
			fmt.Print(i)
			if i == 26 {
				break
			}
			letter <- true
		}
		done <- true
	}()

	letter <- true
	<-done
}
