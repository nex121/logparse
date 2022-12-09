package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type PrefetchList struct {
	Date string `json:"Date"`
	Time string `json:"Time"`
	Type string `json:"Type"`
	Size string `json:"Size"`
	Name string `json:"Name"`
}

func GetPrefetchList() {
	var pl PrefetchList
	var plArray []PrefetchList
	cmd := exec.Command("cmd", "/c", "dir", "C:\\Windows\\Prefetch")
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
			pl.Date = result3[0]
			pl.Time = result3[1]
			pl.Type = "DIR"
			pl.Name = result3[3]
			plArray = append(plArray, pl)
		} else {
			pl.Date = result3[0]
			pl.Time = result3[1]
			pl.Type = "LNK/FILE"
			pl.Size = result3[2]
			pl.Name = result3[3]
			plArray = append(plArray, pl)
		}
	}

	plJson, _ := json.Marshal(plArray)

	encrypt, _ := utils.EncryptByAes(plJson)
	err = ioutil.WriteFile("./Output/prefetch.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] Prefetch收集成功")
	}
}
