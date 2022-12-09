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

type SystemStartupList struct {
	Caption  string `json:"Caption"`
	Command  string `json:"Command"`
	Location string `json:"Location"`
	User     string `json:"User"`
	UserSID  string `json:"UserSID"`
}

func GetSystemStartupList() {
	var ssl SystemStartupList
	var sslArray []SystemStartupList
	cmd := exec.Command("wmic", "startup", "get", "Caption,Command,Location,User,UserSID", "/format:csv")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))

	}

	result0 := utils.ConvertByte2String(out, "GB18030")

	if strings.Contains(result0, "没有可用实例") {
		ssl.Caption = "null"
		ssl.Command = "null"
		ssl.Location = "null"
		ssl.User = "null"
		ssl.UserSID = "null"
		sslArray = append(sslArray, ssl)
		sslJson, _ := json.Marshal(sslArray)
		encrypt, _ := utils.EncryptByAes(sslJson)
		err = ioutil.WriteFile("./Output/SystemStartupList.json", []byte(encrypt), 0777)
		if err != nil {
			fmt.Println("[-] 收集错误", err.Error())
		} else {
			fmt.Println("[+] Program收集成功")
		}
		return
	}
	result1 := strings.Split(result0, "\r\n")
	sliceLen := len(result1) - 1
	result2 := result1[2:sliceLen]

	for _, i := range result2 {
		result3 := strings.Split(i, ",")
		ssl.Caption = result3[1]
		ssl.Command = result3[2]
		ssl.Location = result3[3]
		ssl.User = result3[4]
		ssl.UserSID = result3[5]
		sslArray = append(sslArray, ssl)
	}
	//
	sslJson, _ := json.Marshal(sslArray)
	encrypt, _ := utils.EncryptByAes(sslJson)
	err = ioutil.WriteFile("./Output/SystemStartupList.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] SystemStartup收集成功")
	}
}
