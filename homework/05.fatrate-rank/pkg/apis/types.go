package apis

type PersonalInformation struct {
	Name   string  `json:"name"`
	Sex    string  `json:"sex"`
	Age    int     `json:"age"`
	Tall   float64 `json:"tall"`
	Weight float64 `json:"weight"`
}
