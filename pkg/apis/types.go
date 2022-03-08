package apis

type PersonalInformation struct {
	name   string  `json:"name"` //注意，私有的成员变量在序列化、反序列化时会被忽略
	Name   string  `json:"name"`
	Sex    string  `json:"sex"`
	Age    int     `json:"age"`
	Tall   float64 `json:"tall"`
	Weight float64 `json:"weight"`
}
