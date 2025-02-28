package workerpool

import (
	"fmt"
	"sync"
)

type Executable interface {
	Execute(output chan<- any) error
}

type WorkerPool struct {
	tasks   chan Executable
	errors  chan error
	outputs chan any
	close   chan bool

	wg     sync.WaitGroup
	taskWg sync.WaitGroup
}

func New(workerCount int, taskWorkerRatio int) *WorkerPool {
	wp := &WorkerPool{
		tasks:   make(chan Executable, workerCount*taskWorkerRatio),
		errors:  make(chan error),
		outputs: make(chan any),
		close:   make(chan bool),
		wg:      sync.WaitGroup{},
		taskWg:  sync.WaitGroup{},
	}

	for i := range workerCount {
		wp.wg.Add(1)
		go wp.worker(i)
	}

	return wp
}

func (wp *WorkerPool) Submit(task Executable) {
	wp.tasks <- task
}

func (wp *WorkerPool) Close() {
	wp.taskWg.Wait()
	close(wp.close)
	wp.wg.Wait()
	close(wp.tasks)
}

func (wp *WorkerPool) worker(id int) {
	fmt.Printf("Starting worker #%d\n", id)
	wp.wg.Done()

	for {
		select {
		case task := <-wp.tasks:
			wp.taskWg.Add(1)
			
			fmt.Printf("New task for worker #%d\n", id)
			if err := task.Execute(wp.outputs); err != nil {
				wp.errors <- err
			}
			
			wp.taskWg.Done()
		case <-wp.close:
			return
		}
	}
}
