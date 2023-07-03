package cpu

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

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

func(c *cpu) PrintProcessState(ID  string){
	// Get the process ID of the target process
	pid := 1234 // Replace with the actual process ID

	// Open the process using the process ID
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the process state
	state, err := process.Signal(syscall.Signal(0)).to
	if err != nil {
		log.Fatal(err)
	}

	// Print the process state
	fmt.Println("Process state:", state)
}
