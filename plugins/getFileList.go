package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/dlclark/regexp2"
	"io/fs"
	"io/ioutil"
	"logparseProject/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

type FileList struct {
	Name           string `json:"Name"`
	Type           string `json:"Type"`
	Size           string `json:"Size"`
	CreateTime     string `json:"ModTime"`
	ModTime        string `json:"UserSID"`
	LastAccessTime string `json:"LastAccessTime"`
}

func GetFileList() {
	var fl FileList
	var flArray []FileList
	reg1 := regexp2.MustCompile(`(?<=\').*(?=\')`, 0)
	cmd := exec.Command("cmd", "/c", `%temp%`)
	out, _ := cmd.CombinedOutput()
	result0 := utils.ConvertByte2String(out, "GB18030")
	result00, _ := reg1.FindStringMatch(result0)
	filepath.Walk("C:\\", func(path string, d os.FileInfo, err error) error { //nolint:errcheck
		if err != nil {
			return nil
		} else {
			fileinfo, err1 := os.Stat(path)
			if err1 != nil {
				fl.Size = strconv.FormatInt(fileinfo.Size(), 10)
				fl.CreateTime = "null"
				fl.ModTime = fileinfo.ModTime().Format("2006-01-02 15:04:05")
				fl.LastAccessTime = "null"
				return nil
			}
			fl.Name = path
			if fileinfo.IsDir() {
				fl.Type = "DIR"
			} else {
				fl.Type = "FILE"
			}
			fileSys := fileinfo.Sys().(*syscall.Win32FileAttributeData)
			fl.Size = strconv.FormatInt(fileinfo.Size(), 10)
			fl.CreateTime = utils.SecondToTime(fileSys.CreationTime.Nanoseconds() / 1e9)
			fl.ModTime = utils.SecondToTime(fileSys.LastWriteTime.Nanoseconds() / 1e9)
			fl.LastAccessTime = utils.SecondToTime(fileSys.LastWriteTime.Nanoseconds() / 1e9)
			flArray = append(flArray, fl)
		}
		if d.IsDir() && strings.Count(path, string(os.PathSeparator)) > 1 || d.Name() == "Users" {
			return fs.SkipDir
		}
		return nil
	})

	filepath.Walk(result00.String(), func(path string, d os.FileInfo, err error) error { //nolint:errcheck
		if err != nil {
			return nil
		} else {
			fileinfo, err1 := os.Stat(path)
			if err1 != nil {
				fl.Size = strconv.FormatInt(fileinfo.Size(), 10)
				fl.CreateTime = "null"
				fl.ModTime = fileinfo.ModTime().Format("2006-01-02 15:04:05")
				fl.LastAccessTime = "null"
				return nil
			}
			fl.Name = path
			if fileinfo.IsDir() {
				fl.Type = "DIR"
			} else {
				fl.Type = "FILE"
			}
			fileSys := fileinfo.Sys().(*syscall.Win32FileAttributeData)
			fl.Size = strconv.FormatInt(fileinfo.Size(), 10)
			fl.CreateTime = utils.SecondToTime(fileSys.CreationTime.Nanoseconds() / 1e9)
			fl.ModTime = utils.SecondToTime(fileSys.LastWriteTime.Nanoseconds() / 1e9)
			fl.LastAccessTime = utils.SecondToTime(fileSys.LastWriteTime.Nanoseconds() / 1e9)
			flArray = append(flArray, fl)
		}
		counts := strings.Count(result0, string(os.PathSeparator))
		if d.IsDir() && strings.Count(path, string(os.PathSeparator)) > counts+4 || d.Name() == "Users" {
			return fs.SkipDir
		}
		return nil
	})
	flJson, _ := json.Marshal(flArray)
	encrypt, _ := utils.EncryptByAes(flJson)
	err := ioutil.WriteFile("./Output/FileList.json", []byte(encrypt), 0777)
	if err != nil {
		fmt.Println("[-] FileList收集失败", err.Error())
	} else {
		fmt.Println("[+] FileList收集成功")
	}
}
