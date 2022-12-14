package cmdKit

import (
	"os/exec"
)

// Execute 执行命令（会阻塞直到命令结束）
func Execute(name string, args ...string) ([]byte, error) {
	/*
		!!!:
		(1) exec.Cmd结构体执行时，会处理路径中的空格（e.g. java可执行文件的绝对路径、-Djava.ext.dirs=的路径...）
		(2) 假如自行处理命令行中的路径，反而会导致: 命令执行失败
	*/
	cmd := exec.Command(name, args...)

	return cmd.CombinedOutput()
}

// ExecuteToString 执行命令（会阻塞直到命令结束）
func ExecuteToString(name string, args ...string) (string, error) {
	data, err := Execute(name, args...)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
