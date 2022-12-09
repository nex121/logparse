package common

import (
	"fmt"
	"logparseProject/plugins"
	"logparseProject/utils"
	"os"
)

func Logo() {
	banner := `
        ###    ########  ######## ######## ########  ########
    ## ##   ##     ##    ##    ##       ##     ##    ##
   ##   ##  ##     ##    ##    ##       ##     ##    ##
  ##     ## ########     ##    ######   ########     ##
  ######### ##     ##    ##    ##       ##   ##      ##
  ##     ## ##     ##    ##    ##       ##    ##     ##
  ##     ## ########     ##    ######## ##     ##    ##

                     abtert version: ` + version + `
`
	fmt.Print(banner)
}

func OutputPathExists() {
	_, err := os.Stat("Output")
	_, err = os.Stat("Output/event")
	if os.IsNotExist(err) {
		err1 := os.Mkdir("Output", 0777)
		err1 = os.Mkdir("Output/event", 0777)
		if err1 != nil {
			return
		}
	}
}

func Flag() {
	Logo()
	OutputPathExists()
	plugins.GetAccount()
	plugins.GetArp()
	//plugins.GetClip()
	plugins.GetDnsCatch()
	plugins.GetDriveList()
	plugins.GetEventLog()
	plugins.GetFileSenDir()
	plugins.GetFireWallList()
	plugins.GetHistory()
	plugins.GetHosts()
	plugins.GetKbList()
	plugins.GetNetWorksList()
	plugins.GetOthers()
	plugins.GetPipeName()
	//plugins.GetPrefetchList()
	plugins.GetProcess()
	plugins.GetProgramList()
	plugins.GetRecentFile()
	plugins.GetRegeditList()
	plugins.GetRouteList()
	plugins.GetSchTasksList()
	plugins.GetServicesList()
	plugins.GetSharesList()
	plugins.GetStartUpFile()
	plugins.GetSystemStartupList()
	plugins.GetUserTempFile()
	plugins.GetWmiObjectList()
	plugins.GetFileList()
	plugins.GetSystemInfoList()

	utils.Zip("Output/", "Output.zip")
}
