package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
)

type DnsCatchList struct {
	DnsCatch string `json:"DnsCatch"`
}

func GetDnsCatch() {
	var dcl DnsCatchList

	cmd := exec.Command("ipconfig", "/displaydns")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
	}
	result0 := utils.ConvertByte2String(out, "GB18030")
	dcl.DnsCatch = result0
	dclJson, _ := json.Marshal(dcl)
	encrypt, _ := utils.EncryptByAes(dclJson)
	err = ioutil.WriteFile("./Output/dnsCatch.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] DnsCatch收集成功")
	}
}
