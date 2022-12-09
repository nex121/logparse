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

type NetWorksList struct {
	Protocol        string `json:"Protocol"`
	LocalAddress    string `json:"LocalAddress"`
	ExternalAddress string `json:"ExternalAddress"`
	Status          string `json:"Status"`
	PID             string `json:"PID"`
}

func GetNetWorksList() {
	var nwl NetWorksList
	var nwlArray []NetWorksList
	cmd := exec.Command("netstat", "-ano")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
	}

	result0 := utils.ConvertByte2String(out, "GB18030")
	result1 := strings.Split(result0, "\r\n")
	sliceLen := len(result1) - 1
	result2 := result1[4:sliceLen]

	for _, i := range result2 {
		result3 := strings.Fields(i)
		if len(result3) == 5 {
			nwl.Protocol = result3[0]
			nwl.LocalAddress = result3[1]
			nwl.ExternalAddress = result3[2]
			nwl.Status = result3[3]
			nwl.PID = result3[4]
			nwlArray = append(nwlArray, nwl)
		} else {
			nwl.Protocol = result3[0]
			nwl.LocalAddress = result3[1]
			nwl.ExternalAddress = result3[2]
			nwl.Status = "null"
			nwl.PID = result3[3]
			nwlArray = append(nwlArray, nwl)
		}
	}
	//
	nwlJson, _ := json.Marshal(nwlArray)
	encrypt, _ := utils.EncryptByAes(nwlJson)
	err = ioutil.WriteFile("./Output/networks.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] Networks收集成功")
	}
}
