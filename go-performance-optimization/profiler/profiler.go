package profiler

import (
	"log"
	"net/http"
	_ "net/http/pprof" // Import the pprof package for profiling
)

// StartProfiler starts the pprof profiler
func StartProfiler() {
	go func() {
		log.Println("Starting profiler on http://localhost:6060/debug/pprof/")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("Error starting profiler: %v", err)
		}
	}()
}
