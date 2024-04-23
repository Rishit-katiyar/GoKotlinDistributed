package main

import (
	"sync"
	"time"
)

// Coordinator coordinates task distribution
type Coordinator struct {
	Workers []*Worker
	TaskQueue chan Task
	WorkerStatus map[int]bool
	mutex sync.Mutex
}

// NewCoordinator creates a new coordinator
func NewCoordinator(numWorkers int) *Coordinator {
	coordinator := &Coordinator{
		TaskQueue: make(chan Task),
		WorkerStatus: make(map[int]bool),
	}
	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(i)
		coordinator.Workers = append(coordinator.Workers, worker)
		coordinator.WorkerStatus[i] = true
		worker.Start()
		go coordinator.monitorWorkerStatus(worker)
	}
	go coordinator.dispatchTasks()
	return coordinator
}

// DispatchTasks dispatches tasks from the queue to available workers
func (c *Coordinator) dispatchTasks() {
	for {
		select {
		case task := <-c.TaskQueue:
			c.mutex.Lock()
			workerID := task.ID % len(c.Workers)
			if c.WorkerStatus[workerID] {
				c.Workers[workerID].TaskCh <- task
			}
			c.mutex.Unlock()
		}
	}
}

// MonitorWorkerStatus monitors the status of a worker
func (c *Coordinator) monitorWorkerStatus(worker *Worker) {
	for {
		select {
		case <-time.After(1 * time.Minute):
			c.mutex.Lock()
			workerID := worker.ID
			lastActivityTime := worker.StartTime
			if time.Since(lastActivityTime) > 2*time.Minute {
				c.WorkerStatus[workerID] = false
			} else {
				c.WorkerStatus[workerID] = true
			}
			c.mutex.Unlock()
		}
	}
}

// DistributeTask distributes task to workers
func (c *Coordinator) DistributeTask(task Task) {
	c.TaskQueue <- task
}

// StopWorkers stops all workers
func (c *Coordinator) StopWorkers() {
	for _, worker := range c.Workers {
		worker.Stop()
	}
}



