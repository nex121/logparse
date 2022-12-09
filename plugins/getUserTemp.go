package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type UserTempFile struct {
	Date string `json:"Date"`
	Time string `json:"Time"`
	Type string `json:"Type"`
	Size string `json:"Size"`
	Name string `json:"Name"`
}

func GetUserTempFile() {
	var ut UserTempFile
	var utArray []UserTempFile
	cmd := exec.Command("cmd", "/c", "dir", "%temp%")
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
			ut.Date = result3[0]
			ut.Time = result3[1]
			ut.Type = "DIR"
			ut.Name = result3[3]
			utArray = append(utArray, ut)
		} else {
			ut.Date = result3[0]
			ut.Time = result3[1]
			ut.Type = "LNK/FILE"
			ut.Size = result3[2]
			ut.Name = result3[3]
			utArray = append(utArray, ut)
		}
	}

	utJson, _ := json.Marshal(utArray)
	encrypt, _ := utils.EncryptByAes(utJson)
	err = ioutil.WriteFile("./Output/userTemp.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] UserTemp收集成功")
	}
}
