//AcceptPause,AcceptStop,Caption,CreationClassName,DelayedAutoStart,Description,DesktopInteract,DisplayName,ErrorControl,Name,PathName,ProcessId,ServiceType,Started,StartMode,StartName,State,Status,SystemCreationClassName,SystemName

package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type ServicesList struct {
	Caption           string `json:"Caption"`
	CreationClassName string `json:"CreationClassName"`
	PathName          string `json:"PathName"`
	ProcessId         string `json:"ProcessId"`
	ServiceType       string `json:"ServiceType"`
	Started           string `json:"Started"`
	StartMode         string `json:"StartMode"`
	StartName         string `json:"StartName"`
	State             string `json:"State"`
	Status            string `json:"Status"`
}

func GetServicesList() {
	var sl ServicesList
	var slArray []ServicesList
	cmd := exec.Command("wmic", "service", "get", "Caption,CreationClassName,PathName,ProcessId,ServiceType,Started,StartMode,StartName,State,Status", "/format:csv")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))

	}

	result0 := utils.ConvertByte2String(out, "GB18030")
	result1 := strings.Split(result0, "\r\n")
	sliceLen := len(result1) - 2
	result2 := result1[2:sliceLen]

	for _, i := range result2 {
		result3 := strings.Split(i, ",")
		sl.Caption = result3[1]
		sl.CreationClassName = result3[2]
		sl.PathName = result3[3]
		sl.ProcessId = result3[4]
		sl.ServiceType = result3[5]
		sl.Started = result3[6]
		sl.StartMode = result3[7]
		sl.StartName = result3[8]
		sl.State = result3[9]
		sl.Status = result3[10]
		slArray = append(slArray, sl)
	}
	//
	slJson, _ := json.Marshal(slArray)
	encrypt, _ := utils.EncryptByAes(slJson)
	err = ioutil.WriteFile("./Output/Services.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] Service收集失败", err.Error())
	} else {
		fmt.Println("[+] Service收集成功")
	}
}
