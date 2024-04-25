package main

import (
	"golang.org/x/xerrors"
	"io"
	"net/http"
	"os"
)

func downloadFile(localPath string, remotePath string) error {
	file, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	rsp, err := http.Get(remotePath)
	defer func() {
		_ = rsp.Body.Close()
	}()
	if err != nil {
		return err
	}

	if rsp.StatusCode != 200 {
		return xerrors.Errorf("down file error code: %d", rsp.StatusCode)
	}
	_, err = io.Copy(file, rsp.Body)
	return err
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}
