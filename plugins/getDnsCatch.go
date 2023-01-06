package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type DnsCatchList struct {
	RecordName   string `json:"RecordName"`
	RecordType   string `json:"RecordType"`
	SurvivalTime string `json:"SurvivalTime"`
	DataLength   string `json:"DataLength"`
	Part         string `json:"Part"`
	Records      string `json:"Records"`
}

func GetDnsCatch() {
	var dcl DnsCatchList
	var dclArray []DnsCatchList
	cmd := exec.Command("ipconfig", "/displaydns")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
	}
	result0 := utils.ConvertByte2String(out, "GB18030")
	result1 := strings.Split(result0, "\r\n\r\n\r\n")

	for _, i := range result1 {
		if strings.Contains(i, "---------------------") {
			result2 := strings.Split(i, "----------------------------------------\r\n")
			if strings.Contains(result2[0], "Windows IP 配置") {
				result2[0] = strings.ReplaceAll(result2[0], "Windows IP 配置", "")
			}
			if strings.Contains(result2[1], " . . . :") {
				result3 := strings.Split(result2[1], "\r\n")
				dcl.RecordName = strings.TrimSpace(strings.Split(result3[0], ":")[1])
				dcl.RecordType = strings.TrimSpace(strings.Split(result3[1], ":")[1])
				dcl.SurvivalTime = strings.TrimSpace(strings.Split(result3[2], ":")[1])
				dcl.DataLength = strings.TrimSpace(strings.Split(result3[3], ":")[1])
				dcl.Part = strings.TrimSpace(strings.Split(result3[4], ":")[1])
				dcl.Records = strings.TrimSpace(strings.Split(result3[5], ":")[1])
			}
		}
		dclArray = append(dclArray, dcl)
	}
	dclJson, _ := json.Marshal(dclArray)

	encrypt, _ := utils.EncryptByAes(dclJson)
	//err = ioutil.WriteFile("./Output/dnsCatch.json", dclJson, 0777)
	err = ioutil.WriteFile("./Output/dnsCatch.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] DnsCatch收集失败", err.Error())
	} else {
		fmt.Println("[+] DnsCatch收集成功")
	}
}
