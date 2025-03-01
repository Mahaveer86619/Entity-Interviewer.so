package services

import (
    "time"
    "github.com/shirou/gopsutil/host"
    "github.com/shirou/gopsutil/process"
)

func GetUptime() (time.Duration, error) {
    uptime, err := host.Uptime()
    if err != nil {
        return 0, err
    }
    return time.Duration(uptime) * time.Second, nil
}

func GetProcessInfo() (process.Process, error) {
    pid := int32(1) // Assuming PID 1 for the main process
    proc, err := process.NewProcess(pid)
    if err != nil {
        return process.Process{}, err
    }
    return *proc, nil
}