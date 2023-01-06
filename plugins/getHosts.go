package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type HostList struct {
	Ip     string `json:"Ip"`
	Domain string `json:"Domain"`
}

func GetHosts() {
	var hl HostList
	var hlArray []HostList
	cmd := exec.Command("more", "C:\\Windows\\System32\\drivers\\etc\\hosts")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
	}
	result0 := utils.ConvertByte2String(out, "GB18030")
	result1 := strings.Split(result0, "\r\n")
	for _, i := range result1 {
		if !strings.Contains(i, "#") && len(strings.Fields(i)) != 0 {
			hl.Ip = strings.Fields(i)[0]
			hl.Domain = strings.Fields(i)[1]
			hlArray = append(hlArray, hl)
		}
	}
	hlJson, _ := json.Marshal(hlArray)
	encrypt, _ := utils.EncryptByAes(hlJson)
	err = ioutil.WriteFile("./Output/hosts.json", []byte(encrypt), 0777)

	if err != nil {
		fmt.Println("[-] Hosts收集失败", err.Error())
	} else {
		fmt.Println("[+] Hosts收集成功")
	}
}
