//AcceptPause,AcceptStop,Caption,CreationClassName,DelayedAutoStart,Description,DesktopInteract,DisplayName,ErrorControl,Name,PathName,ProcessId,ServiceType,Started,StartMode,StartName,State,Status,SystemCreationClassName,SystemName

package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
	"strings"
)

type SchTasksList struct {
	TaskName        string `json:"TaskName"`
	NextRunTime     string `json:"NextRunTime"`
	Mode            string `json:"Mode"`
	LoginStatus     string `json:"LoginStatus"`
	LastRunTime     string `json:"LastRunTime"`
	LastResult      string `json:"LastResult"`
	Creator         string `json:"Creator"`
	RunTask         string `json:"RunTask"`
	StartFrom       string `json:"StartFrom"`
	ExplanatoryNote string `json:"ExplanatoryNote"`
	SchTaskStatus   string `json:"SchTaskStatus"`
	WhoRun          string `json:"WhoRun"`
	StartTime       string `json:"StartTime"`
	StartDate       string `json:"StartDate"`
	StopDate        string `json:"StopDate"`
}

func GetSchTasksList() {
	var stl SchTasksList
	var stlArray []SchTasksList
	cmd := exec.Command("SCHTASKS", "/Query", "/FO", "csv", "/V", "/NH")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))

	}

	result0 := utils.ConvertByte2String(out, "GB18030")

	result1 := strings.Split(result0, "\r\n")
	sliceLen := len(result1) - 1
	result2 := result1[:sliceLen]
	for _, i := range result2 {
		result3 := strings.Split(i, "\",\"")
		stl.TaskName = result3[1]
		stl.NextRunTime = result3[2]
		stl.Mode = result3[3]
		stl.LoginStatus = result3[4]
		stl.LastRunTime = result3[5]
		stl.LastResult = result3[6]
		stl.Creator = result3[7]
		stl.RunTask = result3[8]
		stl.StartFrom = result3[9]
		stl.ExplanatoryNote = result3[10]
		stl.SchTaskStatus = result3[11]
		stl.WhoRun = result3[14]
		stl.StartTime = result3[19]
		stl.StartDate = result3[20]
		stl.StopDate = result3[21]
		stlArray = append(stlArray, stl)
	}

	stlJson, _ := json.Marshal(stlArray)
	encrypt, _ := utils.EncryptByAes(stlJson)
	err = ioutil.WriteFile("./Output/SchTaskList.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] 收集错误", err.Error())
	} else {
		fmt.Println("[+] Schtasks收集成功")
	}
}
