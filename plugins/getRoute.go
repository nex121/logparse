//AcceptPause,AcceptStop,Caption,CreationClassName,DelayedAutoStart,Description,DesktopInteract,DisplayName,ErrorControl,Name,PathName,ProcessId,ServiceType,Started,StartMode,StartName,State,Status,SystemCreationClassName,SystemName

package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
)

type RouteList struct {
	Route string `json:"Route"`
}

func GetRouteList() {
	var rl RouteList
	cmd := exec.Command("route", "print")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))

	}

	result0 := utils.ConvertByte2String(out, "GB18030")
	rl.Route = result0
	rlJson, _ := json.Marshal(rl)
	encrypt, _ := utils.EncryptByAes(rlJson)
	err = ioutil.WriteFile("./Output/route.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] Route收集成功")
	}
}
