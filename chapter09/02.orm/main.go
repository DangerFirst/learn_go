package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:123qwe@tcp(localhost)/testdb"))
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func createNewPerson(conn *gorm.DB) error {
	resp := conn.Create(&PersonalInformation{
		Name:   "小黑20220322",
		Sex:    "男",
		Age:    24,
		Tall:   1.77,
		Weight: 63,
	})
	if err := resp.Error; err != nil {
		fmt.Println("创建人***时失败：", err)
		return err
	}
	fmt.Println("创建***人成功")
	return nil
}

func ormSelect(conn *gorm.DB) {
	outPut := make([]*PersonalInformation, 0)
	resp := conn.Where(&PersonalInformation{Name: "小黑"}).Find(&outPut)
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return
	}
	data, _ := json.Marshal(outPut)
	fmt.Printf("查询结果：%+v\n", string(data))
}

func ormSelectWithInaccurateQuery(conn *gorm.DB) {
	outPut := make([]*PersonalInformation, 0)
	resp := conn.Where("tall>?", 1.7).Find(&outPut)
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return
	}
	data, _ := json.Marshal(outPut)
	fmt.Printf("查询结果：%+v\n", string(data))
}

func updateExistingPerson(conn *gorm.DB) error {
	resp := conn.Updates(&PersonalInformation{
		Id:     6,
		Name:   "小黑20220322",
		Sex:    "男",
		Age:    100,
		Tall:   1.77,
		Weight: 63,
	})
	if err := resp.Error; err != nil {
		fmt.Println("更新人***时失败：", err)
		return err
	}
	fmt.Println("更新***人成功")
	return nil
}

func updateExistingPersonSelectFields(conn *gorm.DB) error {
	p := &PersonalInformation{
		Id:     6,
		Name:   "小黑202203221",
		Sex:    "男",
		Age:    19,
		Tall:   1.78,
		Weight: 66,
	}
	resp := conn.Model(p).Select("name", "tall").Updates(p)
	if err := resp.Error; err != nil {
		fmt.Println("更新人***时失败：", err)
		return err
	}
	fmt.Println("更新***人成功")
	return nil
}

func deletePerson(conn *gorm.DB) error {
	p := &PersonalInformation{
		Id: 6,
	}
	resp := conn.Delete(p)
	if err := resp.Error; err != nil {
		fmt.Println("删除人***时失败：", err)
		return err
	}
	fmt.Println("删除***人成功")
	return nil
}

func main() {
	conn := connectDb()
	fmt.Println(conn)
	//if err:=createNewPerson(conn);err!=nil{
	//	//todo handle error
	//}
	//ormSelect(conn)
	//ormSelectWithInaccurateQuery(conn)
	//updateExistingPerson(conn)
	//updateExistingPersonSelectFields(conn)
	deletePerson(conn)
}
