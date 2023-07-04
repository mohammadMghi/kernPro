package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"strconv"
	"strings"

	cpu "github.com/mohammadMghi/kernPro/cpu"
)

func main() {

	flagString := flag.String("type", "", "This flag is for determin which part of application you wnat to use")
	pid := flag.String("pid", "-1", "Specify a string flag")
 

	flag.Parse()

	cpu := cpu.NewCpu()

	if(*flagString == "process"){
		readProccess()
		return
	}
	
	if *flagString == "cpu"{
		cpu.CpuUtilization()
 
		return
	}

	if *flagString == "cpu_heavy"{
		
		cpu.PrintHeavyProcesses()
 
		return
	}

	if *flagString == "process_state" && *pid != "-1"{
		
		cpu.PrintProcessState(*pid)

		return
	} 

	if *flagString == "process_child" && *pid != "-1"{
		num, err := strconv.Atoi(*pid)
		if err != nil{
			log.Fatal(err.Error())

		}
		cpu.PrintProcessChild(num)
	}

	if *flagString == "sys_process" {
		cpu.PrintSysProcess()
	}


	fmt.Printf("Please set a flag like k=cpu for checking cpu utilization ...");

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