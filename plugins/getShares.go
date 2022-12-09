//AccessMask  AllowMaximum  Caption   Description  InstallDate  MaximumAllowed  Name    Path        Status  Type

package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type SharesList struct {
	AccessMask   string `json:"AccessMask"`
	AllowMaximum string `json:"AllowMaximum"`
	Caption      string `json:"Caption"`
	Name         string `json:"Name"`
	Path         string `json:"Path"`
	Status       string `json:"Status"`
	Type         string `json:"Type"`
}

func GetSharesList() {
	var sl SharesList
	var slArray []SharesList
	cmd := exec.Command("wmic", "share", "get", "AccessMask,AllowMaximum,Caption,Name,Path,Status,Type", "/format:csv")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))

	}

	result0 := utils.ConvertByte2String(out, "GB18030")
	result1 := strings.Split(result0, "\r\n")
	sliceLen := len(result1) - 1
	result2 := result1[2:sliceLen]

	for _, i := range result2 {
		result3 := strings.Split(i, ",")
		sl.AccessMask = result3[1]
		sl.AllowMaximum = result3[2]
		sl.Caption = result3[3]
		sl.Name = result3[4]
		sl.Path = result3[5]
		sl.Status = result3[6]
		sl.Type = result3[7]
		slArray = append(slArray, sl)
	}

	slJson, _ := json.Marshal(slArray)
	encrypt, _ := utils.EncryptByAes(slJson)
	err = ioutil.WriteFile("./Output/shareList.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] Shares收集成功")
	}
}
