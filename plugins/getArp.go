package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/dlclark/regexp2"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"regexp"
	"strings"
)

type ArpList struct {
	Interface       string `json:"Interface"`
	InternetAddress string `json:"InternetAddress"`
	MacAddress      string `json:"MacAddress"`
	Type            string `json:"Type"`
}

func GetArp() {
	var al ArpList
	var alArray []ArpList
	reg1 := regexp2.MustCompile(`(?<=:).*(?=---)`, 0)
	reg2 := regexp.MustCompile(`接口.*\r\n|Internet.*\r\n`)
	cmd := exec.Command("arp", "-a")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
	}
	result0 := utils.ConvertByte2String(out, "GB18030")

	result1 := strings.Split(result0, "\r\n\r\n")
	for _, i := range result1 {
		result2 := reg2.ReplaceAll([]byte(i), []byte(""))
		result3, _ := reg1.FindStringMatch(i)

		result4 := strings.Split(string(result2), "\r\n")
		result4Len := len(result4)
		result5 := result4[1 : result4Len-1]
		for _, j := range result5 {
			result6 := strings.Fields(j)
			al.Interface = strings.TrimSpace(result3.String())
			al.InternetAddress = result6[0]
			al.MacAddress = result6[1]
			al.Type = result6[2]
			alArray = append(alArray, al)
		}
	}
	alJson, _ := json.Marshal(alArray)
	encrypt, _ := utils.EncryptByAes(alJson)
	err = ioutil.WriteFile("./Output/arp.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] Arp收集失败", err.Error())
	} else {
		fmt.Println("[+] Arp收集成功")
	}
}
