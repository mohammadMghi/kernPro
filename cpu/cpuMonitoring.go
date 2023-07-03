package cpu

import (
	"bufio"
	"fmt"
	"io/ioutil"

	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"sync"

	"strconv"

	"time"
)


type cpu struct{

}


func NewCpu() *cpu{
	return &cpu{}
}

 

func (c *cpu)getCPUTemperature() (string, error) {
	file, err := os.Open("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	temperature := scanner.Text()

	return temperature, nil
}

func (c *cpu)PrintTempState() {
	temperature, err := c.getCPUTemperature()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	temperatureFloat, err := strconv.ParseFloat(temperature, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	temperatureCelsius := temperatureFloat / 1000.0
	temperatureFormatted := fmt.Sprintf("%.2f°C", temperatureCelsius)

	fmt.Println("CPU Temperature:", temperatureFormatted)
}


func (c *cpu)CpuUtilization() {


	// Create a ticker that ticks every second
	ticker := time.NewTicker(1 * time.Second)

	go func() {

		for range ticker.C{
			// Get the initial CPU usage
			prevTime := time.Now()
			prevCPU := c.getCPUUsage()

			// Wait for a specific duration
			time.Sleep(1 * time.Second)

			// Get the updated CPU usage
			currTime := time.Now()
			currCPU := c.getCPUUsage()

			// Calculate the CPU utilization percentage
			totalTime := currTime.Sub(prevTime).Seconds()
			cpuUtilization := 100 * (float64(currCPU-prevCPU) / float64(runtime.NumCPU()) / totalTime)

			fmt.Printf("CPU Utilization: %.2f%%\n", cpuUtilization)


			c.PrintTempState()
		}
	
	}()	

	select {}

 

}

// Returns the current CPU time as reported by the Go runtime
func (c *cpu) getCPUUsage() int64 {
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return int64(s.Sys)
}


func (c *cpu)HeavyFunction(wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate a heavy computation
	for i := 0; i < 1000000000; i++ {
		_ = i * i
	}
}

func (c *cpu)PrintHeavyProcesses() {
	// Execute the 'ps aux' command to retrieve information about all processes
	cmd := exec.Command("ps", "aux")

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// Split the output into lines
	lines := strings.Split(string(output), "\n")

	// Print the header
	fmt.Println(lines[0])

	// Print the heavy processes
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) != "" {
			fmt.Println(line)
		}
	}
}

func(c *cpu) PrintProcessState(pid  string){
 
 

	// Execute the 'ps -p <pid> -o pid,ppid,cmd' command to get process information
	cmd := exec.Command("ps", "-p", pid, "-o", "pid,ppid,cmd")

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// Split the output into lines
	lines := strings.Split(string(output), "\n")

	// Extract the process information from the output
	if len(lines) > 1 {
		fields := strings.Fields(lines[1])
		if len(fields) == 3 {
			processID := fields[0]
			parentProcessID := fields[1]
			command := fields[2]
			fmt.Println("Process ID:", processID)
			fmt.Println("Parent Process ID:", parentProcessID)
			fmt.Println("Command:", command)
		}
	} else {
		fmt.Println("Process not found")
	}
}

// Parse the process state from the status file content
func parseProcessState(status string) string {
	lines := strings.Split(status, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "State:") {
			fields := strings.Fields(line)
			if len(fields) > 1 {
				return fields[1]
			}
		}
	}
	return ""
}


func (c *cpu) PrintProcessChild(processId int){
 

	// Get the list of processes in the /proc directory
	processes, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the processes and find the child processes of the parent PID
	childPIDs := make([]int, 0)
	for _, process := range processes {
		if process.IsDir() {
			// Check if the process directory name is a valid PID
			pid, err := strconv.Atoi(process.Name())
			if err == nil && isChildProcess(pid, processId) {
				childPIDs = append(childPIDs, pid)
			}
		}
	}

	// Print the child process IDs
	fmt.Println("Child Processes:")
	for _, childPID := range childPIDs {
		fmt.Println(childPID)
	}
}


// Check if the given PID is a child process of the parent PID
func isChildProcess(pid, parentPID int) bool {
	// Read the stat file of the process to get its parent PID
	statFile := fmt.Sprintf("/proc/%d/stat", pid)
	statData, err := ioutil.ReadFile(statFile)
	if err != nil {
		return false
	}

	// Extract the parent PID from the stat file data
	statFields := strings.Fields(string(statData))
	if len(statFields) > 3 {
		ppid, err := strconv.Atoi(statFields[3])
		if err == nil && ppid == parentPID {
			return true
		}
	}

	return false
}