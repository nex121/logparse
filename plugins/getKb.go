package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type KbList struct {
	Name        string `json:"Name"`
	CSName      string `json:"CSName"`
	Description string `json:"Description"`
	HotFixID    string `json:"HotFixID"`
	InstalledBy string `json:"InstalledBy"`
	InstalledOn string `json:"InstalledOn"`
}

func GetKbList() {
	var kl KbList
	var klArray []KbList
	cmd := exec.Command("wmic", "qfe", "get", "Caption,CSName,Description,HotFixID,InstalledBy,InstalledOn", "/format:csv")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
	}

	result0 := utils.ConvertByte2String(out, "GB18030")
	if strings.Contains(result0, "没有可用实例") {
		kl.Name = "null"
		kl.CSName = "null"
		kl.Description = "null"
		kl.HotFixID = "null"
		kl.InstalledBy = "null"
		kl.InstalledOn = "null"
		klArray = append(klArray, kl)
		klJson, _ := json.Marshal(klArray)
		encrypt, _ := utils.EncryptByAes(klJson)
		err = ioutil.WriteFile("./Output/kbList.json", []byte(encrypt), 0777)
		if err != nil {
			fmt.Println("[-] 收集错误", err.Error())
		} else {
			fmt.Println("[+] Kb收集成功")
		}
		return
	}
	result1 := strings.Split(result0, "\r\n")
	sliceLen := len(result1) - 1
	result2 := result1[2:sliceLen]

	for _, i := range result2 {
		result3 := strings.Split(i, ",")
		kl.Name = result3[0]
		kl.CSName = result3[1]
		kl.Description = result3[2]
		kl.HotFixID = result3[3]
		kl.InstalledBy = result3[4]
		kl.InstalledOn = result3[5]
		klArray = append(klArray, kl)
	}

	klJson, _ := json.Marshal(klArray)
	encrypt, _ := utils.EncryptByAes(klJson)
	err = ioutil.WriteFile("./Output/kbList.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] Kb收集成功")
	}
}
