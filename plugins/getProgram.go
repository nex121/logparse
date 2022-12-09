package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type ProgramList struct {
	Name          string `json:"Name"`
	IdNumber      string `json:"IdNumber"`
	InstallDate   string `json:"InstallDate"`
	InstallSource string `json:"InstallSource"`
	LocalPackage  string `json:"LocalPackage"`
}

func GetProgramList() {
	var pl ProgramList
	var plArray []ProgramList
	cmd := exec.Command("wmic", "product", "get", "Caption,IdentifyingNumber,InstallDate,InstallSource,LocalPackage", "/format:csv")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
	}

	result0 := utils.ConvertByte2String(out, "GB18030")

	if strings.Contains(result0, "没有可用实例") {
		pl.Name = "null"
		pl.IdNumber = "null"
		pl.InstallDate = "null"
		pl.InstallSource = "null"
		pl.LocalPackage = "null"
		plArray = append(plArray, pl)
		klJson, _ := json.Marshal(plArray)
		encrypt, _ := utils.EncryptByAes(klJson)
		err = ioutil.WriteFile("./Output/programList.json", []byte(encrypt), 0777)
		if err != nil {
			fmt.Println("[-] 收集错误", err.Error())
		} else {
			fmt.Println("[+] Program收集成功")
		}
		return
	}

	result1 := strings.Split(result0, "\r\n")
	sliceLen := len(result1) - 1
	result2 := result1[2:sliceLen]

	for _, i := range result2 {
		result3 := strings.Split(i, ",")
		pl.Name = result3[1]
		pl.IdNumber = result3[2]
		pl.InstallDate = result3[3]
		pl.InstallSource = result3[4]
		pl.LocalPackage = result3[5]
		plArray = append(plArray, pl)
	}

	plJson, _ := json.Marshal(plArray)
	encrypt, _ := utils.EncryptByAes(plJson)
	err = ioutil.WriteFile("./Output/programList.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] Program收集成功")
	}
}
