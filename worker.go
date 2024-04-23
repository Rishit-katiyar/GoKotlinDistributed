package main

import (
	"time"
)

// Worker represents a Go worker node
type Worker struct {
	ID           int
	TaskCh       chan Task
	QuitCh       chan struct{}
	StartTime    time.Time
	TotalTasks   int
	CompletedTasks int
}

// NewWorker creates a new worker
func NewWorker(id int) *Worker {
	return &Worker{
		ID:           id,
		TaskCh:       make(chan Task),
		QuitCh:       make(chan struct{}),
		StartTime:    time.Now(),
		TotalTasks:   0,
		CompletedTasks: 0,
	}
}

// Start starts the worker
func (w *Worker) Start() {
	go func() {
		for {
			select {
			case task := <-w.TaskCh:
				w.TotalTasks++
				// Simulate processing by adding "Processed" prefix
				task.Result = "Processed: " + task.Input
				tasksLock.Lock()
				tasks = append(tasks, task)
				tasksLock.Unlock()
				w.CompletedTasks++
			case <-w.QuitCh:
				return
			}
		}
	}()
}

// Stop stops the worker
func (w *Worker) Stop() {
	close(w.QuitCh)
}

// GetWorkerStats returns the statistics of the worker
func (w *Worker) GetWorkerStats() WorkerStats {
	return WorkerStats{
		ID:            w.ID,
		TotalTasks:    w.TotalTasks,
		CompletedTasks: w.CompletedTasks,
		StartTime:     w.StartTime,
		LastActivityTime: time.Now(),
	}
}

// WorkerStats represents statistics of a worker
type WorkerStats struct {
	ID             int
	TotalTasks     int
	CompletedTasks int
	StartTime      time.Time
	LastActivityTime time.Time
}
