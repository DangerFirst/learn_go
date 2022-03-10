package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"learn.go/pkg/apis"
	"log"
)

func main() {
	learnDB, err := getDbConnection()
	defer learnDB.Close()
	err = testDbConnection(err, learnDB)

	queryAllData(err, learnDB)
	insertData(learnDB) //todo handle error
	queryAllData(err, learnDB)
}

func insertData(learnDB *sql.DB) error {
	_, err := learnDB.Exec(fmt.Sprintf("insert into personal_information(name,sex,tall,weight,age) values('%s','%s',%f,%f,%d)",
		"小妹",
		"女",
		1.75,
		55.0,
		18,
	))
	if err != nil {
		fmt.Println("新增数据失败：", err)
		return err
	}
	return nil
}

func queryAllData(err error, learnDB *sql.DB) {
	rows, err := learnDB.Query("select name,age from personal_information")
	if err != nil {
		log.Fatal(err)
	}

	list := &apis.PersonalInformationList{}
	for rows.Next() {
		var name string
		var age int
		rows.Scan(&name, &age) //todo handle error
		if err := rows.Scan(&name, &age); err != nil {
			log.Printf("扫描数据失败，跳过该行：%v", err)
		}
		fmt.Printf("name:%s\tage:%d\n", name, age)
		list.Items = append(list.Items, &apis.PersonalInformationP{
			Name: name,
			Age:  int64(age),
		})
	}
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}

func testDbConnection(err error, learnDB *sql.DB) error {
	if err = learnDB.Ping(); nil != err {
		log.Fatal("DB测试失败：", err)
	}
	fmt.Println("数据库连接成功，可以执行命令")
	return err
}

func getDbConnection() (*sql.DB, error) {
	learnDB, err := sql.Open("mysql", "root:123qwe@tcp(localhost)/testdb")
	if err != nil {
		log.Fatal("链接数据库失败:", err)
	}
	return learnDB, err
}
