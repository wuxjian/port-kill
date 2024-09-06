package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	process_infos, err := getProcessInfoByPort("8085")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(process_infos)
}

func getProcessInfoByPort(port string) ([]ProcessInfo, error) {
	cmd := exec.Command("cmd", "/c", "netstat -ano | findstr "+port)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := string(out)
	// result 按换行分割
	lines := strings.Split(result, "\r\n")
	var processInfo []ProcessInfo
	for _, line := range lines {
		// line 按照正则\tab分割
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		// fields[0] 协议
		// fields[1] 内部地址
		// fields[2] 外部地址
		// fields[3] 状态
		// fields[4] 进程号
		//
		fmt.Println(fields[0], fields[1], fields[2], fields[3], fields[4])
		processInfo = append(processInfo,
			ProcessInfo{Protocol: fields[0],
				LocalAddress:  fields[1],
				RemoteAddress: fields[2],
				State:         fields[3],
				Pid:           fields[4]})
	}
	return processInfo, nil
}

type ProcessInfo struct {
	Protocol      string `json:"protocol"`
	LocalAddress  string `json:"local_address"`
	RemoteAddress string `json:"remote_address"`
	State         string `json:"state"`
	Pid           string `json:"pid"`
}
