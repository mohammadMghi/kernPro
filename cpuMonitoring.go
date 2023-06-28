package main


import (
	"fmt"
	"runtime"
	"time"
)

func cpuUtilization() {
	// Get the initial CPU usage
	prevTime := time.Now()
	prevCPU := getCPUUsage()

	// Wait for a specific duration
	time.Sleep(1 * time.Second)

	// Get the updated CPU usage
	currTime := time.Now()
	currCPU := getCPUUsage()

	// Calculate the CPU utilization percentage
	totalTime := currTime.Sub(prevTime).Seconds()
	cpuUtilization := 100 * (float64(currCPU-prevCPU) / float64(runtime.NumCPU()) / totalTime)

	fmt.Printf("CPU Utilization: %.2f%%\n", cpuUtilization)
}

// Returns the current CPU time as reported by the Go runtime
func getCPUUsage() int64 {
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return int64(s.Sys)
}

