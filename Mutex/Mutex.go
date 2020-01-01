package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁，可以有多个读
var RWLock = sync.RWMutex{}

// 互斥锁
var Lock = sync.Mutex{}
var wg = sync.WaitGroup{}

func readTask(number int) {
	// 读锁可以有多个读（CPU数）
	defer wg.Done()
	RWLock.RLock()
	defer RWLock.RUnlock()
	fmt.Printf("[%d]Start Read.\n", number)
	time.Sleep(5 * time.Second)
	fmt.Printf("[%d]End Read.\n", number)
}

func task(number int) {
	// 只能一个任务在执行
	defer wg.Done()
	Lock.Lock()
	defer Lock.Unlock()
	fmt.Printf("[%d]Start.\n", number)
	time.Sleep(5 * time.Second)
	fmt.Printf("[%d]End.\n", number)

}

func main() {
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go readTask(i)
	}

	wg.Add(2)
	for i := 0; i < 2; i++ {
		go task(i)
	}
	wg.Wait()
}
