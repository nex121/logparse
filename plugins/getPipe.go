package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type PipeName struct {
	Date string `json:"Date"`
	Time string `json:"Time"`
	Type string `json:"Type"`
	Size string `json:"Size"`
	Name string `json:"Name"`
}

func GetPipeName() {
	var pn PipeName
	var pnArray []PipeName
	cmd := exec.Command("cmd", "/c", "dir", "\\\\.\\pipe\\\\")
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
			pn.Date = result3[0]
			pn.Time = result3[1]
			pn.Type = "DIR"
			pn.Name = result3[3]
			pnArray = append(pnArray, pn)
		} else {
			pn.Date = result3[0]
			pn.Time = result3[1]
			pn.Type = "LNK/FILE"
			pn.Size = result3[2]
			pn.Name = result3[3]
			pnArray = append(pnArray, pn)
		}
	}

	pnJson, _ := json.Marshal(pnArray)
	encrypt, _ := utils.EncryptByAes(pnJson)
	err = ioutil.WriteFile("./Output/pipeName.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] Pipe收集失败", err.Error())
	} else {
		fmt.Println("[+] Pipe收集成功")
	}
}
