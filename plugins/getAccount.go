package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type UserAccount struct {
	AccountType     string `json:"AccountType"`
	Caption         string `json:"Caption"`
	Domain          string `json:"Domain"`
	LocalAccount    string `json:"LocalAccount"`
	PasswordExpires string `json:"PasswordExpires"`
	SID             string `json:"SID"`
	Status          string `json:"Status"`
}

func GetAccount() {
	var ua UserAccount
	var uaArray []UserAccount
	cmd := exec.Command("wmic", "useraccount", "get", "AccountType,Caption,Domain,LocalAccount,PasswordExpires,SID,Status")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
	}
	result := strings.Split(string(out), "\r\n")[1:]

	for _, i := range result {
		result1 := strings.Fields(i)
		if len(result1) > 1 {

			ua.AccountType = result1[0]
			ua.Caption = result1[1]
			ua.Domain = result1[2]
			ua.LocalAccount = result1[3]
			ua.PasswordExpires = result1[4]
			ua.SID = result1[5]
			ua.Status = result1[6]
			uaArray = append(uaArray, ua)
		}
	}
	uaJson, _ := json.Marshal(uaArray)
	encrypt, _ := utils.EncryptByAes(uaJson)
	err = ioutil.WriteFile("./Output/account.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] Account收集成功")
	}
}
