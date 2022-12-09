package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type FileSenDir struct {
	Date string `json:"Date"`
	Time string `json:"Time"`
	Type string `json:"Type"`
	Size string `json:"Size"`
	Name string `json:"Name"`
}

func GetFileSenDir() {
	var rf RecentFile
	var rfArray []RecentFile
	cmd := exec.Command("cmd", "/c", "dir", "%USERPROFILE%\\AppData\\Roaming\\Microsoft\\Windows\\Recent")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
	}

	result0 := utils.ConvertByte2String(out, "GB18030")

	result1 := strings.Split(result0, "\r\n")
	sliceLen := len(result1) - 3
	result2 := result1[5:sliceLen]
	for _, i := range result2 {

		result3 := strings.Fields(i)
		if strings.Contains(i, "<DIR>") {
			rf.Date = result3[0]
			rf.Time = result3[1]
			rf.Type = "DIR"
			rf.Name = result3[3]
			rfArray = append(rfArray, rf)
		} else {
			rf.Date = result3[0]
			rf.Time = result3[1]
			rf.Type = "LNK"
			rf.Size = result3[2]
			rf.Name = result3[3]
			rfArray = append(rfArray, rf)
		}
	}

	rfJson, _ := json.Marshal(rfArray)
	encrypt, _ := utils.EncryptByAes(rfJson)
	err = ioutil.WriteFile("./Output/recent.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] FileSensitive收集成功")
	}
}
