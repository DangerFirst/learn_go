package apis

//
//type PersonalInformation struct {
//	name          string  `json:"name"` //注意，私有的成员变量在序列化、反序列化时会被忽略
//	Name          string  `json:"name"`
//	Sex           string  `json:"sex"`
//	Age           int     `json:"age"`
//	Tall          float64 `json:"tall"`
//	Weight        float64 `json:"weight"`
//	state         interface{}
//	sizeCache     interface{}
//	unknownFields interface{}
//}
//
//
//type PersonalInformationFatRate struct {
//	Name    string
//	FatRate float64
//}
//
//type PersonalRank struct {
//	Name       string
//	Sex        string
//	RankNumber int
//	FatRate    float64
//}

func (*PersonalInformation) TableName() string {
	return "personal_information"
}

func (*PersonalShow) TableName() string {
	return "personal_show"
}
