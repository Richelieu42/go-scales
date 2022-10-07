package runtimeKit

import (
	"bytes"
	"gitee.com/richelieu042/go-scales/src/core/errorKit"
	"gitee.com/richelieu042/go-scales/src/core/strKit"
	"github.com/shirou/gopsutil/v3/cpu"
)

// GetCpuId 获取cpu的id
/*
PS: cpu.Info() 目前仅支持Windows环境，不支持：Mac环境（M1）、Linux环境.
*/
func GetCpuId() (string, error) {
	stats, err := cpu.Info()
	if err != nil {
		return "", err
	}
	if len(stats) == 0 {
		return "", errorKit.Simple("length of stats is 0")
	}

	buffer := bytes.Buffer{}
	for _, stat := range stats {
		if strKit.IsNotEmpty(stat.PhysicalID) {
			if buffer.Len() > 0 {
				// 多个cpu id间的分隔符
				buffer.WriteString("-")
			}
			buffer.WriteString(stat.PhysicalID)
		}
	}
	if buffer.Len() == 0 {
		return "", errorKit.Simple("length of buffer is 0")
	}
	return buffer.String(), nil
}
