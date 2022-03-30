package main

import (
	"fmt"
	"gorm.io/gorm"
	"learn.go/chapter11/02.practice/frinterface"
	"learn.go/pkg/apis"
	"log"
)

var _ frinterface.ServerInterface = &dbRank{}
var _ frinterface.RankInitInterface = &dbRank{}

func NewDbRank(conn *gorm.DB,
	embedRank frinterface.ServerInterface) frinterface.ServerInterface {
	if conn == nil {
		log.Fatal("数据库连接为空")
	}
	if embedRank == nil {
		log.Fatal("内存排行榜为空")
	}
	return &dbRank{
		conn:      conn,
		embedRank: embedRank,
	}
}

type dbRank struct {
	conn      *gorm.DB
	embedRank frinterface.ServerInterface
}

func (d dbRank) Init() error {
	outPut := make([]*apis.PersonalInformation, 0)
	resp := d.conn.Find(&outPut)
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return resp.Error
	}
	for _, item := range outPut {
		if _, err := d.embedRank.UpdatePersonalInformation(item); err != nil {
			log.Fatalf("初始化%s时失败：%v\n", item.Name, err)
		}
	}
	return nil
}

func (d dbRank) RegisterPersonalInformation(pi *apis.PersonalInformation) error {
	resp := d.conn.Create(pi)
	if err := resp.Error; err != nil {
		// 注意：不同企业对log有要求，比如：必须带上某个ID。log时使用公司各自的log框架
		// e.g. https://github.com/sirupsen/logrus/
		fmt.Printf("创建%s时失败：%v\n", pi.Name, err)
		return err
	}
	fmt.Printf("创建%s成功\n", pi.Name)
	_ = d.embedRank.RegisterPersonalInformation(pi) //todo handle error
	return nil
}

func (d dbRank) UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error) {
	resp := d.conn.Updates(pi)
	if err := resp.Error; err != nil {
		fmt.Printf("更新%s时失败：%v\n", pi.Name, err)
		return nil, err
	}
	fmt.Printf("更新%s成功\n", pi.Name)
	return d.embedRank.UpdatePersonalInformation(pi)
}

func (d dbRank) GetFatRate(name string) (*apis.PersonalRank, error) {
	return d.embedRank.GetFatRate(name)
}

func (d dbRank) GetTop() ([]*apis.PersonalRank, error) {
	return d.embedRank.GetTop()
}
