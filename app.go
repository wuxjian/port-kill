package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// 执行cmd方法，并返回结果
func (a *App) GetProcessInfoByPort(port string) ([]ProcessInfo, error) {
	cmd := exec.Command("cmd", "/c", "netstat -ano | findstr "+port)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("get port %s failed %v", port, err)
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

func (a *App) Kill(ports string) error {
	portArray := strings.Split(ports, ",")
	for _, port := range portArray {
		err := kill(port)
		if err != nil {
			fmt.Printf("kill port %s failed %v", port, err)
			return err
		}
	}
	return nil
}

func kill(port string) error {
	cmd := exec.Command("cmd", "/c", fmt.Sprintf("taskkill /PID %s /F", port))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err := cmd.Output()
	return err
}

type ProcessInfo struct {
	Protocol      string `json:"protocol"`
	LocalAddress  string `json:"local_address"`
	RemoteAddress string `json:"remote_address"`
	State         string `json:"state"`
	Pid           string `json:"pid"`
}
