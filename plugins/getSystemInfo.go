package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"io/ioutil"
	"logparseProject/utils"
	"strconv"
)

type SystemInfoList struct {
	CpuModelName         string `json:"CpuModelName"`
	CpuHz                string `json:"CpuHz"`
	Hostname             string `json:"Hostname"`
	Os                   string `json:"Os"`
	Platform             string `json:"Platform"`
	PlatformFamily       string `json:"PlatformFamily"`
	PlatformVersion      string `json:"PlatformVersion"`
	KernelVersion        string `json:"kernelVersion"`
	KernelArch           string `json:"KernelArch"`
	VirtualizationSystem string `json:"VirtualizationSystem"`
	TotalPhysicalMemory  string `json:"TotalPhysicalMemory"`
	MemoryAvailable      string `json:"MemoryAvailable"`
	MemoryUsed           string `json:"MemoryUsed"`
}

func GetSystemInfoList() {
	var sil SystemInfoList
	cpus, _ := cpu.Info()
	hosts, _ := host.Info()
	mems, _ := mem.VirtualMemory()
	sil.CpuModelName = cpus[0].ModelName
	sil.CpuHz = strconv.Itoa(int(cpus[0].Mhz))
	sil.Hostname = hosts.Hostname
	sil.Os = hosts.OS
	sil.Platform = hosts.Platform
	sil.PlatformFamily = hosts.PlatformFamily
	sil.PlatformVersion = hosts.PlatformVersion
	sil.KernelVersion = hosts.KernelVersion
	sil.KernelArch = hosts.KernelArch
	sil.VirtualizationSystem = hosts.VirtualizationSystem
	sil.TotalPhysicalMemory = strconv.Itoa(int(mems.Total))
	sil.MemoryUsed = strconv.Itoa(int(mems.Used))
	sil.MemoryAvailable = strconv.Itoa(int(mems.Available))

	silJson, _ := json.Marshal(sil)
	encrypt, _ := utils.EncryptByAes(silJson)
	err := ioutil.WriteFile("./Output/SystemInfo.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] SystemInfo收集失败", err.Error())
	} else {
		fmt.Println("[+] SystemInfo收集成功")
	}
}
