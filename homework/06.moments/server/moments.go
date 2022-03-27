package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"moments/frinterface"
	"moments/pkg/apis"
	"time"
)

type MomentsItem struct {
	Name            string
	ShowDescription string
	ShowTime        string
}

type Moments struct {
	items []MomentsItem
}

func (m Moments) Publish(ps *apis.PersonalShow) error {
	return nil
}

func (m Moments) Delete(ps *apis.PersonalShow) error {
	return nil
}

func (m Moments) Moments() ([]*apis.GoMoments, error) {
	return nil, nil
}

type dbMoments struct {
	conn         *gorm.DB
	embedMoments frinterface.MomentsServerInterface
}

func (d dbMoments) Publish(ps *apis.PersonalShow) error {
	ps = &apis.PersonalShow{
		PersonalId:      ps.PersonalId,
		PersonalName:    ps.PersonalName,
		ShowDescription: ps.ShowDescription,
		ShowTime:        time.Now().UTC().String(),
		ByTimeAge:       ps.ByTimeAge,
		ByTimeTall:      ps.ByTimeTall,
		ByTimeWeight:    ps.ByTimeWeight,
		ByTimeFatRate:   ps.ByTimeFatRate,
		IsDeleted:       0,
	}

	resp := d.conn.Create(ps)
	if err := resp.Error; err != nil {
		fmt.Printf("%s发布朋友圈失败：%v\n", ps.PersonalName, err)
		return err
	}
	fmt.Printf("%s发布成功\n", ps.PersonalName)
	if err := d.embedMoments.Publish(ps); err != nil {
		fmt.Printf("%s发布朋友圈失败：%v\n", ps.PersonalName, err)
		return err
	}
	return nil
}

func (d dbMoments) Delete(ps *apis.PersonalShow) error {
	resp := d.conn.Model(&apis.PersonalShow{}).Where("personal_name=?", ps.PersonalName).Update("is_deleted", 1)
	if err := resp.Error; err != nil {
		fmt.Printf("删除%s的朋友圈状态失败：%v\n", ps.PersonalName, err)
		return err
	}
	fmt.Printf("删除%s的朋友圈状态成功\n", ps.PersonalName)
	return nil
}

func (d dbMoments) Moments() ([]*apis.GoMoments, error) {
	var tables []apis.PersonalShow
	getMoments := d.conn.Where("is_deleted=?", 0).Order("id desc").Find(&tables)
	if err := getMoments.Error; err != nil {
		fmt.Printf("刷新朋友圈失败：%v\n", err)
		return nil, err
	}
	out := make([]*apis.GoMoments, 0)
	for i, _ := range tables {
		item := tables[i]
		out = append(out, &apis.GoMoments{
			Name:            item.PersonalName,
			ShowDescription: item.ShowDescription,
			ShowTime:        item.ShowTime,
		})
	}
	fmt.Printf("刷新朋友圈成功\n")
	return out, nil
}

var _ frinterface.MomentsServerInterface = &dbMoments{}

func NewDbMoments(conn *gorm.DB,
	embedMoments frinterface.MomentsServerInterface) frinterface.MomentsServerInterface {
	if conn == nil {
		log.Fatal("数据库连接为空")
	}
	if embedMoments == nil {
		log.Fatal("内存排行榜为空")
	}
	return &dbMoments{
		conn:         conn,
		embedMoments: embedMoments,
	}
}

func NewMoments() *Moments {
	return &Moments{
		items: make([]MomentsItem, 0, 100),
	}
}
