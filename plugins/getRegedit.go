package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
)

type RegeditList struct {
	CurrentUserRun      string `json:"CurrentUserRun"`
	CurrentUserRunOnce  string `json:"CurrentUserRunOnce"`
	LocalMachineRun     string `json:"LocalMachineRun"`
	LocalMachineRunOnce string `json:"LocalMachineRunOnce"`
	UAC                 string `json:"UAC"`
}

func GetRegeditList() {
	var rl RegeditList
	cmd := exec.Command("reg", "query", "HKEY_CURRENT_USER\\Software\\Microsoft\\Windows\\CurrentVersion\\Run")
	cmd1 := exec.Command("reg", "query", "HKEY_CURRENT_USER\\Software\\Microsoft\\Windows\\CurrentVersion\\RunOnce")
	cmd2 := exec.Command("reg", "query", "HKEY_LOCAL_MACHINE\\Software\\Microsoft\\Windows\\CurrentVersion\\Run")
	cmd3 := exec.Command("reg", "query", "HKEY_CURRENT_USER\\Software\\Microsoft\\Windows\\CurrentVersion\\RunOnce")
	cmd4 := exec.Command("reg", "query", "HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Policies\\System")

	out, err := cmd.CombinedOutput()
	out1, err := cmd1.CombinedOutput()
	out2, err := cmd2.CombinedOutput()
	out3, err := cmd3.CombinedOutput()
	out4, err := cmd4.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))

	}
	result0 := utils.ConvertByte2String(out, "GB18030")
	result1 := utils.ConvertByte2String(out1, "GB18030")
	result2 := utils.ConvertByte2String(out2, "GB18030")
	result3 := utils.ConvertByte2String(out3, "GB18030")
	result4 := utils.ConvertByte2String(out4, "GB18030")

	rl.CurrentUserRun = result0
	rl.CurrentUserRunOnce = result1
	rl.LocalMachineRun = result2
	rl.CurrentUserRunOnce = result3
	rl.UAC = result4

	rlJson, _ := json.Marshal(rl)
	encrypt, _ := utils.EncryptByAes(rlJson)
	err = ioutil.WriteFile("./Output/regedit.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] Regedit收集失败", err.Error())
	} else {
		fmt.Println("[+] Regedit收集成功")
	}
}
