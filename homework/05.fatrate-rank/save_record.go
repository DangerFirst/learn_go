package main

import (
	"encoding/json"
	"fatrate-rank/pkg/apis"
	"fmt"
	"log"
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
	if err := r.writeFileWithAppendJson(data); err != nil {
		log.Println("写入json时出错：", err)
		return err
	}
	return nil
}

func (r *record) writeFileWithAppendJson(data []byte) error {
	file, err := os.OpenFile(r.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
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
