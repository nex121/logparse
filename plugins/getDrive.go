package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type DriveList struct {
	BytesPerSector          string `json:"BytesPerSector"`
	Capabilities            string `json:"Capabilities"`
	CapabilityDescriptions  string `json:"CapabilityDescriptions"`
	Caption                 string `json:"Caption"`
	ConfigManagerUserConfig string `json:"ConfigManagerUserConfig"`
	CreationClassName       string `json:"CreationClassName"`
	Description             string `json:"Description"`
	DeviceID                string `json:"DeviceID"`
	FirmwareRevision        string `json:"FirmwareRevision"`
	Index                   string `json:"Index"`
	InstallDate             string `json:"InstallDate"`
}

func GetDriveList() {
	var dl DriveList
	var dlArray []DriveList
	cmd := exec.Command("wmic", "diskdrive", "get", "BytesPerSector,Capabilities,CapabilityDescriptions,Caption,ConfigManagerUserConfig,CreationClassName,Description,DeviceID,FirmwareRevision,InstallDate,Index", "/format:csv")
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
		dl.BytesPerSector = result3[1]
		dl.Capabilities = result3[2]
		dl.CapabilityDescriptions = result3[3]
		dl.Caption = result3[4]
		dl.ConfigManagerUserConfig = result3[5]
		dl.CreationClassName = result3[6]
		dl.Description = result3[7]
		dl.DeviceID = result3[8]
		dl.FirmwareRevision = result3[9]
		dl.Index = result3[10]
		dl.InstallDate = result3[11]
		dlArray = append(dlArray, dl)
	}

	dlJson, _ := json.Marshal(dlArray)
	encrypt, _ := utils.EncryptByAes(dlJson)
	err = ioutil.WriteFile("./Output/driveList.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] Drive收集成功")
	}
}
