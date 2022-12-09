package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

//File copies a single file from src to dst
func File(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

// Dir copies a whole directory recursively
func Dir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = Dir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = File(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

//func CopyFile(src, dst string) (err error) {
//	sfi, err := os.Stat(src)
//	if err != nil {
//		return
//	}
//	if !sfi.Mode().IsRegular() {
//		// cannot copy non-regular files (e.g., directories,
//		// symlinks, devices, etc.)
//		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
//	}
//	dfi, err := os.Stat(dst)
//	if err != nil {
//		if !os.IsNotExist(err) {
//			return
//		}
//	} else {
//		if !(dfi.Mode().IsRegular()) {
//			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
//		}
//		if os.SameFile(sfi, dfi) {
//			return
//		}
//	}
//	if err = os.Link(src, dst); err == nil {
//		return
//	}
//	err = copyFileContents(src, dst)
//	return
//}
//
//func copyFileContents(src, dst string) (err error) {
//	in, err := os.Open(src)
//	if err != nil {
//		return
//	}
//	defer in.Close()
//	out, err := os.Create(dst)
//	if err != nil {
//		return
//	}
//	defer func() {
//		cerr := out.Close()
//		if err == nil {
//			err = cerr
//		}
//	}()
//	if _, err = io.Copy(out, in); err != nil {
//		return
//	}
//	err = out.Sync()
//	return
//}
