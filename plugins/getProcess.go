package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type TaskList struct {
	CSName         string `json:"Name"`
	CreationDate   string `json:"CreationDate"`
	Caption        string `json:"Caption"`
	ExecutablePath string `json:"ExecutablePath"`
	ProcessId      string `json:"ProcessId"`
}

func GetProcess() {
	var tl TaskList
	var tlArray []TaskList
	cmd := exec.Command("wmic", "process", "get", "CSName,CreationDate,Caption,ExecutablePath,ProcessId", "/format:csv")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
	}
	result0 := utils.ConvertByte2String(out, "GB18030")
	result1 := strings.Split(result0, "\r\n")[2:]
	sliceLen := len(result1) - 1
	result2 := result1[2:sliceLen]
	for _, i := range result2 {
		result3 := strings.Split(i, ",")
		if len(result3) == 0 {
			continue
		}
		tl.CSName = result3[1]
		tl.CreationDate = result3[2]
		tl.Caption = result3[3]
		tl.ExecutablePath = result3[4]
		tl.ProcessId = strings.Trim(result3[5], "\r")
		tlArray = append(tlArray, tl)
	}

	tlJson, _ := json.Marshal(tlArray)
	encrypt, _ := utils.EncryptByAes(tlJson)
	err = ioutil.WriteFile("./Output/process.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] Process收集失败", err.Error())
	} else {
		fmt.Println("[+] Process收集成功")
	}
}
