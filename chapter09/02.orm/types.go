package main

type PersonalInformation struct {
	Id     int64   `json:"id" gorm:"primaryKey;column:id"`
	Name   string  `json:"name" gorm:"column:name"`
	Sex    string  `json:"sex" gorm:"column:sex"`
	Age    int     `json:"age" gorm:"column:age"`
	Tall   float64 `json:"tall" gorm:"column:tall"`
	Weight float64 `json:"weight" gorm:"column:weight"`
}

func (*PersonalInformation) TableName() string {
	return "personal_information"
}
