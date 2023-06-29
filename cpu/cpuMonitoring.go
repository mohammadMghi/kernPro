package cpu

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"strconv"

	"time"

	"google.golang.org/appengine/runtime"
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
	// Get the current Goroutine count
	numGoroutines := runtime.NumGoroutine()

	// Create a slice to store Goroutine IDs
	goroutineIDs := make([]int64, numGoroutines)

	// Capture the Goroutine IDs
	n := runtime.Getgoroutines(&goroutineIDs[0], true)

	// Print the Goroutine IDs and their CPU usage
	for i := 0; i < n; i++ {
		goid := goroutineIDs[i]
		cpuUsage := runtime.CPUProfilePercent(goid)

		fmt.Printf("Goroutine ID: %d, CPU Usage: %.2f%%\n", goid, cpuUsage)
	}
}
