package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logparseProject/utils"
	"os/exec"
)

type Clip struct {
	ClipContent string `json:"ClipContent"`
}

func GetClip() {
	//powershell Get-Clipboard
	var cl Clip
	cmd := exec.Command("powershell", "Get-Clipboard")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
	}
	result := string(out)
	cl.ClipContent = result
	clJson, _ := json.Marshal(cl)
	encrypt, _ := utils.EncryptByAes(clJson)
	err = ioutil.WriteFile("./Output/clip.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("Clip收集失败", err.Error())
	} else {
		fmt.Println("Clip收集成功")
	}
}
