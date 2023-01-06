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

type RouteList struct {
	NetworkObjectives string `json:"NetworkObjectives"`
	Netmask           string `json:"Netmask"`
	Gateway           string `json:"Gateway"`
	Interface         string `json:"Interface"`
	LeapPoints        string `json:"LeapPoints"`
}

func GetRouteList() {
	var rl RouteList
	var rlArray []RouteList
	cmd := exec.Command("route", "print")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))

	}

	result0 := utils.ConvertByte2String(out, "GB18030")
	result1 := strings.Split(result0, "===========================================================================\r\n")
	result2 := strings.Split(result1[3], "\r\n")
	result2Len := len(result2)
	result3 := result2[2 : result2Len-1]
	for _, i := range result3 {
		result4 := strings.Fields(i)
		rl.NetworkObjectives = result4[0]
		rl.Netmask = result4[1]
		rl.Gateway = result4[2]
		rl.Interface = result4[3]
		rl.LeapPoints = result4[4]
		rlArray = append(rlArray, rl)
	}

	rlJson, _ := json.Marshal(rlArray)
	encrypt, _ := utils.EncryptByAes(rlJson)
	//err = ioutil.WriteFile("./Output/route.json", rlJson, 0777)
	err = ioutil.WriteFile("./Output/route.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] Route收集失败", err.Error())
	} else {
		fmt.Println("[+] Route收集成功")
	}
}
