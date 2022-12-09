package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
)

type WmiObjectList struct {
	WmiEventFilter          string `json:"WmiEventFilter"`
	WmiEventConsumer        string `json:"WmiEventConsumer"`
	WmiEventConsumerBinding string `json:"WmiEventConsumerBinding"`
}

func GetWmiObjectList() {
	var wol WmiObjectList
	cmd := exec.Command("powershell", "Get-WMIObject", "-Namespace", "root\\Subscription", "-Class", "__EventFilter")
	cmd1 := exec.Command("powershell", "Get-WMIObject", "-Namespace", "root\\Subscription", "-Class", "CommandLineEventConsumer")
	cmd2 := exec.Command("powershell", "Get-WMIObject", "-Namespace", "root\\Subscription", "-Class", "__FilterToConsumerBinding")
	out, err := cmd.CombinedOutput()
	out1, err := cmd1.CombinedOutput()
	out2, err := cmd2.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
	}
	result0 := utils.ConvertByte2String(out, "GB18030")
	result1 := utils.ConvertByte2String(out1, "GB18030")
	result2 := utils.ConvertByte2String(out2, "GB18030")
	wol.WmiEventFilter = result0
	wol.WmiEventConsumer = result1
	wol.WmiEventConsumerBinding = result2

	wolJson, _ := json.Marshal(wol)
	encrypt, _ := utils.EncryptByAes(wolJson)
	err = ioutil.WriteFile("./Output/wmiObject.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] WmiObject收集成功")
	}
}
