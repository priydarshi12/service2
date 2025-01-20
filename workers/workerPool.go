package workers

import (
	"fmt"
	"sync"
)

var (
	TaskQueue = make(chan func(), 10) // Buffered task queue
	wg        sync.WaitGroup          // WaitGroup for workers
)

func StartWorkerPool(workerCount int) {
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range TaskQueue {
				fmt.Printf("Worker %d processing task\n", workerID)
				task()
			}
		}(i)
	}
}

func WaitForWorkers() {
	wg.Wait()
}
