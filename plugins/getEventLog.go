package plugins

import (
	"fmt"
	"logparseProject/utils"
)

func GetEventLog() {
	logAddress := "C:\\Windows\\System32\\winevt\\Logs\\Security.evtx"
	logAddress1 := "C:\\Windows\\System32\\winevt\\Logs\\System.evtx"
	err1 := utils.File(logAddress, "Output/event/Security.evtx")
	err2 := utils.File(logAddress1, "Output/event/System.evtx")
	if err1 != nil {
		fmt.Println("[-] eventLog收集失败")
	}
	if err2 != nil {
		fmt.Println("[-] eventLog收集失败")
	}
	fmt.Println("[+] eventLog收集成功")
}
