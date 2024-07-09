package common

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

var nvidiaFetchInfo NvidiaFetchInfo

type NvidiaFetchInfo struct {
	MemoryTotal    string
	MemoryUsed     string
	UtilizationGPU string
	GPUName        string
}

func GetNvidiaFetchInfo() NvidiaFetchInfo {
	return NvidiaFetchInfo{}
}

func FetchNvidia() {
	cmd := exec.Command("nvidia-smi", "--query-gpu=index,name,memory.total,memory.used,utilization.gpu", "--format=csv")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating stdout pipe:", err)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		index := fields[0]
		name := fields[1]
		memoryTotal := fields[2]
		memoryUsed := fields[3]
		utilizationGPU := fields[4]

		nvidiaFetchInfo = NvidiaFetchInfo{
			MemoryTotal:    memoryTotal,
			MemoryUsed:     memoryUsed,
			UtilizationGPU: utilizationGPU,
			GPUName:        name,
		}

		fmt.Printf("GPU Index: %s, Name: %s, Memory Total: %s, Memory Used: %s, Utilization GPU: %s\n", index, name, memoryTotal, memoryUsed, utilizationGPU)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning:", err)
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for command to finish:", err)
		return
	}
}
