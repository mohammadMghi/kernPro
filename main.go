package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {

	flagString := flag.String("k", "", "Specify a string flag")

	flag.Parse()

	if(*flagString == "process"){
		readProccess()
		return
	}
	
	if *flagString == "cpu"{
		cpuUtilization()
		return
	}

	


	fmt.Printf("Please set a flag like k=cpu for checking cpu utilization ...");

}

func cpuUtilization() {


	// Create a ticker that ticks every second
	ticker := time.NewTicker(1 * time.Second)

	go func() {

		for range ticker.C{
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
	
	}()	

	select {}

 

}

// Returns the current CPU time as reported by the Go runtime
func getCPUUsage() int64 {
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return int64(s.Sys)
}



func readProccess(){

		// Read the contents of the "/proc" directory
		procDir := "/proc"
		files, err := ioutil.ReadDir(procDir)
		if err != nil {
			fmt.Println("Failed to read /proc directory:", err)
			os.Exit(1)
		}
	
		// Iterate over the directory entries
		for _, file := range files {
			// Check if the entry is a directory with a numeric name (representing a process)
			if file.IsDir() {
				pid, err := strconv.Atoi(file.Name())
				if err != nil {
					continue
				}
	
				// Read the process status file to get information about the process
				statusFilePath := fmt.Sprintf("%s/%d/status", procDir, pid)
				status, err := ioutil.ReadFile(statusFilePath)
				if err != nil {
					continue
				}
	
				// Extract the process name from the status file
				name := extractProcessName(status)
	
				// Print the process ID and name
				fmt.Printf("PID: %d, Name: %s\n", pid, name)
			}
		}
}

// Extracts the process name from the status file contents
func extractProcessName(status []byte) string {
	lines := strings.Split(string(status), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Name:") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				return fields[1]
			}
			break
		}
	}
	return ""
}