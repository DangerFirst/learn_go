package apis

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

func TestMarshalJson(t *testing.T) {
	personalInformation := PersonalInformation{
		name:   `"小"黑"`,
		Sex:    "男",
		Age:    19,
		Tall:   1.77,
		Weight: 65,
	}
	fmt.Printf("%+v\n", personalInformation)

	data, err := json.Marshal(personalInformation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshal的(原生)结果是：", data)
	fmt.Println("marshal的(string)结果是：", string(data))

}

func TestUnmarshalJson(t *testing.T) {
	data := `{"name":"\"小\"黑\"","sex":"男","age":19,"tall":1.77,"weight":65}`
	personalInformation := PersonalInformation{}
	json.Unmarshal([]byte(data), &personalInformation)
	fmt.Println(personalInformation)
}

func TestMarshalYaml(t *testing.T) {
	personalInformation := PersonalInformation{
		name:   `"小"黑"`,
		Sex:    "男",
		Age:    19,
		Tall:   1.77,
		Weight: 65,
	}
	fmt.Printf("%+v\n", personalInformation)

	data, err := yaml.Marshal(personalInformation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshal的(原生)结果是：", data)
	fmt.Println("marshal的(string)结果是：", string(data))

}

func TestUnmarshalJsonYaml(t *testing.T) {
	data := `{"name":"\"小\"黑\"","sex":"男","age":19,"tall":1.77,"weight":65}`
	personalInformation := PersonalInformation{}
	json.Unmarshal([]byte(data), &personalInformation)
	fmt.Println(personalInformation)
}
