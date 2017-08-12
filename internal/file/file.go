package file

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/udhos/equalfile"
)

// MsgOk is when there is no problem - proceed
// MsgNotAnImage is when the file appear not to be an image
// EXIST is when the destination path exist an it seems to be the same file
const (
	MsgOk         string = "Ok              "
	MsgNotAnImage        = "Not an Image    "
)

// CopyFile copies the contents from src to dst atomically.
// If dst does not exist, CopyFile creates it with permissions perm.
// If the copy fails, CopyFile aborts and dst is preserved.
func CopyFile(dst, src string, perm os.FileMode) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	tmp, err := ioutil.TempFile(filepath.Dir(dst), "")
	if err != nil {
		return err
	}
	_, err = io.Copy(tmp, in)
	if err != nil {
		tmp.Close()
		os.Remove(tmp.Name())
		return err
	}
	if err = tmp.Close(); err != nil {
		os.Remove(tmp.Name())
		return err
	}
	if err = os.Chmod(tmp.Name(), perm); err != nil {
		os.Remove(tmp.Name())
		return err
	}
	return os.Rename(tmp.Name(), dst)
}

// MoveFiles move a file from src to dest
// It check if the file exist before and abort if so
func MoveFiles(src, dst string) {
	filepath.Walk(src, moveFile(src, dst))
}

func moveFile(srcPath, dstPath string) func(src string, f os.FileInfo, err error) error {
	return func(srcFile string, f os.FileInfo, err error) error {
		_, filename := path.Split(srcFile)

		var msg string
		var dstFile string

		if dstDatePath, imageErr := DateStringFromImage(srcFile); imageErr == nil {
			dstFilePath := filepath.Join(dstPath, dstDatePath)
			dstFile = path.Join(dstFilePath, filename)
			if _, moveErr := canMoveFile(srcFile, dstFile); moveErr == nil {
				msg = MsgOk
				CheckOrCreatePath(dstFilePath)
				os.Rename(srcFile, dstFile)
			} else {
				msg = fmt.Sprintf("%v", moveErr)
			}
		} else {
			msg = imageErr.Error()
		}

		println(msg, " ", srcFile, " to ", dstFile)
		return nil
	}
}

// checkFile verify if file exist
func canMoveFile(src, dst string) (bool, error) {
	_, dstErr := os.Stat(dst)
	if os.IsExist(dstErr) {
		return true, nil
	}

	srcStat, srcErr := os.Stat(src)
	if os.IsExist(srcErr) || srcStat.IsDir() {
		return false, nil
	}

	if sameFile(src, dst) {
		// println("file are the same ", src, " ", dst)
		return false, errors.New("EXIST  ")
	}

	return true, nil
}

func sameFile(src, dst string) bool {
	options := equalfile.Options{}
	cmp := equalfile.New(nil, options)
	equal, err := cmp.CompareFile(src, dst)
	if err != nil || !equal {
		// fmt.Printf("equal(%s,%s): error: %v\n", src, dst, err)
		return false
	}

	return true
}
