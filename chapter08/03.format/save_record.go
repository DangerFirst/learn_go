package main

import (
	"encoding/json"
	"fmt"
	"learn.go/pkg/apis"
	"os"
)

func NewRecord(filePath string) *record {
	return &record{
		filePath: filePath,
	}
}

type record struct {
	filePath string
}

func (r *record) savePersonalInformation(pi *apis.PersonalInformation) error {
	data, err := json.Marshal(pi)
	if err != nil {
		fmt.Println("marshal出错：", err)
		return err
	}
	return r.writeFileWithAppend(data)
}

func (r *record) writeFileWithAppend(data []byte) error {
	file, err := os.OpenFile(r.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) //linux file permission settings
	if err != nil {
		fmt.Println("无法打开文件", r.filePath, "错误信息为：", err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.Write(append(data, '\n'))
	if err != nil {
		return err
	}
	return nil
}
