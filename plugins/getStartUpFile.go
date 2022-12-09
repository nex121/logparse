package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type StartUpFile struct {
	Date string `json:"Date"`
	Time string `json:"Time"`
	Type string `json:"Type"`
	Size string `json:"Size"`
	Name string `json:"Name"`
}

func GetStartUpFile() {
	var sf RecentFile
	var sfArray []RecentFile
	cmd := exec.Command("cmd", "/c", "dir", "%PROGRAMDATA%\\Microsoft\\Windows\\Start Menu\\Programs\\Startup")
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
			sf.Date = result3[0]
			sf.Time = result3[1]
			sf.Type = "DIR"
			sf.Name = result3[3]
			sfArray = append(sfArray, sf)
		} else {
			sf.Date = result3[0]
			sf.Time = result3[1]
			sf.Type = "LNK/FILE"
			sf.Size = result3[2]
			sf.Name = result3[3]
			sfArray = append(sfArray, sf)
		}
	}

	sfJson, _ := json.Marshal(sfArray)
	encrypt, _ := utils.EncryptByAes(sfJson)
	err = ioutil.WriteFile("./Output/startUp.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] StartupFile收集成功")
	}
}
