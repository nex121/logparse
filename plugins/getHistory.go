package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type History struct {
	PsHistory []string `json:"PsHistory"`
}

func GetHistory() {
	var psHistory []string
	var his History
	_, err := os.Stat("%USERPROFILE%\\AppData\\Roaming\\Microsoft\\Windows\\PowerShell\\PSReadLine\\ConsoleHost_history.txt")
	if os.IsNotExist(err) {
		err1 := os.Mkdir("%USERPROFILE%\\AppData\\Roaming\\Microsoft\\Windows\\PowerShell\\PSReadLine\\", 0777)
		_, err1 = os.OpenFile("%USERPROFILE%\\AppData\\Roaming\\Microsoft\\Windows\\PowerShell\\PSReadLine\\ConsoleHost_history.txt", os.O_APPEND|os.O_CREATE, 0644)
		if err1 != nil {
			return
		}
	}
	cmd := exec.Command("type", "%USERPROFILE%\\AppData\\Roaming\\Microsoft\\Windows\\PowerShell\\PSReadLine\\ConsoleHost_history.txt")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
	}
	result := strings.Split(string(out), "\r\n")

	for _, i := range result {
		psHistory = append(psHistory, i)
	}
	his.PsHistory = psHistory
	hisJson, _ := json.Marshal(his)

	err = ioutil.WriteFile("./Output/history.json", hisJson, 0777)
	if err != nil {
		fmt.Println("[-] History收集失败", err.Error())
	} else {
		fmt.Println("[+] History收集成功")
	}
}
