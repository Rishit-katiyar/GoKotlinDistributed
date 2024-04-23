package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// TaskHandler handles task requests
func TaskHandler(coordinator *Coordinator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse incoming JSON request
		var task Task
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if err := json.Unmarshal(body, &task); err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		// Distribute task to coordinator
		coordinator.DistributeTask(task)

		// Respond with success
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Task received and distributed to workers")
	}
}

// ResultsHandler handles task result requests
func ResultsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasksLock.Lock()
		defer tasksLock.Unlock()

		// Marshal tasks slice to JSON
		response, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, "Failed to marshal task results", http.StatusInternalServerError)
			return
		}

		// Respond with task results
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

// StopHandler handles stopping the workers
func StopHandler(coordinator *Coordinator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		coordinator.StopWorkers()
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Workers stopped successfully")
	}
}

// WorkerStatsHandler handles worker statistics requests
func WorkerStatsHandler(coordinator *Coordinator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		workerStats := make(map[int]WorkerStats)
		for _, worker := range coordinator.Workers {
			workerStats[worker.ID] = worker.GetWorkerStats()
		}

		// Marshal worker stats to JSON
		response, err := json.Marshal(workerStats)
		if err != nil {
			http.Error(w, "Failed to marshal worker stats", http.StatusInternalServerError)
			return
		}

		// Respond with worker stats
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
