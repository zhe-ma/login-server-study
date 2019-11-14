package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func GetVersion(c *gin.Context) {
	c.String(http.StatusOK, "version1.0")
}

func GetComputerInfo(c *gin.Context) {
	// Get Disk info.
	diskName := "C:\\"
	diskStat, _ := disk.Usage(diskName)
	usedSpace := float32(diskStat.Used) / GB
	totalSpace := float32(diskStat.Total) / GB
	freePercent := float32(diskStat.Free) / float32(diskStat.Total) * 100
	diskStr := fmt.Sprintf("Disk (%s) info:\n\tTotal space %.1fGB\n\tUsed space %.1fGB\n\tFree percent %.1f%%\n",
		diskName, totalSpace, usedSpace, freePercent)

	// Get CPU info.
	cpuPhysicalCores, _ := cpu.Counts(false)
	cpuLogicalCores, _ := cpu.Counts(true)
	cpuStr := fmt.Sprintf("CPU cores:\n\tPhysical %d\n\tLogical %d\n", cpuPhysicalCores, cpuLogicalCores)

	// Get memory info.
	memoryStat, _ := mem.VirtualMemory()
	usedSpace = float32(memoryStat.Used) / GB
	totalSpace = float32(memoryStat.Total) / GB
	freePercent = float32(memoryStat.Available) / float32(memoryStat.Total) * 100
	memeryStr := fmt.Sprintf("Memory info:\n\tTotal space %.2fGB\n\tUsed space %.2fGB\n\tFree percent %.1f%%\n",
		totalSpace, usedSpace, freePercent)

	str := diskStr + cpuStr + memeryStr

	c.String(http.StatusOK, str)
}
