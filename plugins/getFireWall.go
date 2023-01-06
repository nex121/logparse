package plugins

//netsh advfirewall firewall show rule name=all dir=in type=dynamic
//netsh advfirewall firewall show rule name=all dir=out type=dynamic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"regexp"
	"strings"
)

type FireWallList struct {
	RuleName          string `json:"RuleName"`
	Status            string `json:"Status"`
	Direction         string `json:"Direction"`
	ConfigurationFile string `json:"ConfigurationFile"`
	Grouping          string `json:"Grouping"`
	LocalIp           string `json:"LocalIp"`
	RemoteIp          string `json:"RemoteIp"`
	Protocol          string `json:"Protocol"`
	LocalPort         string `json:"LocalPort"`
	RemotePort        string `json:"RemotePort"`
	EdgeTraversal     string `json:"EdgeTraversal"`
	Operation         string `json:"Operation"`
}

func GetFireWallList() {
	var fwl FireWallList
	var fwlArray []FireWallList
	reg1 := regexp.MustCompile(`-----.*\r\n`)
	cmd := exec.Command("netsh", "advfirewall", "firewall", "show", "rule", "name=all", "type=dynamic")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fwl.RuleName = "null"
		fwl.Status = "null"
		fwl.Direction = "null"
		fwl.ConfigurationFile = "null"
		fwl.Grouping = "null"
		fwl.LocalIp = "null"
		fwl.RemoteIp = "null"
		fwl.Protocol = "null"
		fwl.LocalPort = "null"
		fwl.RemotePort = "null"
		fwl.EdgeTraversal = "null"
		fwl.Operation = "null"
		fwlArray = append(fwlArray, fwl)
		fwlJson, _ := json.Marshal(fwlArray)
		encrypt, _ := utils.EncryptByAes(fwlJson)
		err = ioutil.WriteFile("./Output/Firewall.json", []byte(encrypt), 0777)
		if err != nil {
			fmt.Println("[-] 收集错误", err.Error())
		} else {
			fmt.Println("[+] Firewall收集成功")
		}
		return
	}

	result0 := utils.ConvertByte2String(out, "GB18030")
	result1 := reg1.ReplaceAll([]byte(result0), []byte(""))
	result4 := strings.Split(string(result1), "\r\n\r\n")
	result4Len := len(result4)
	result5 := strings.Split(string(result1), "\r\n\r\n")[:result4Len-2]
	for _, i := range result5 {
		result6 := strings.Split(i, "\r\n")
		for _, j := range result6 {
			if strings.Contains(j, ":") {
				result7 := strings.Split(j, ":")
				if result7[0] == "规则名称" {
					fwl.RuleName = strings.TrimSpace(result7[1])
				}
				if result7[0] == "已启用" {
					fwl.Status = strings.TrimSpace(result7[1])
				}
				if result7[0] == "方向" {
					fwl.Direction = strings.TrimSpace(result7[1])
				}
				if result7[0] == "配置文件" {
					fwl.ConfigurationFile = strings.TrimSpace(result7[1])
				}
				if result7[0] == "分组" {
					fwl.Grouping = strings.TrimSpace(result7[1])
				}
				if result7[0] == "本地 IP" {
					fwl.LocalIp = strings.TrimSpace(result7[1])
				}
				if result7[0] == "远程 IP" {
					fwl.RemoteIp = strings.TrimSpace(result7[1])
				}
				if result7[0] == "协议" {
					fwl.Protocol = strings.TrimSpace(result7[1])
				}
				if result7[0] == "本地端口" {
					fwl.LocalPort = strings.TrimSpace(result7[1])
				}
				if result7[0] == "远程端口" {
					fwl.RemotePort = strings.TrimSpace(result7[1])
				}
				if result7[0] == "边缘遍历" {
					fwl.EdgeTraversal = strings.TrimSpace(result7[1])
				}
				if result7[0] == "操作" {
					fwl.Operation = strings.TrimSpace(result7[1])
				}
			}
		}
		fwlArray = append(fwlArray, fwl)

	}
	fwlJson, _ := json.Marshal(fwlArray)
	encrypt, _ := utils.EncryptByAes(fwlJson)
	err = ioutil.WriteFile("./Output/Firewall.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] Firewall收集失败", err.Error())
	} else {
		fmt.Println("[+] Firewall收集成功")
	}
}
