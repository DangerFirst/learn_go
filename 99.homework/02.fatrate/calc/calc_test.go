package calc

import (
	"learn.go/99.homework/02.fatrate/input"
	"testing"
)

func TestCalc_BMI(t *testing.T) {
	c := input.Person{
		Weight: 0,
		Tall:   1.7,
	}
	err := Calc{}.BMI(&c)
	if err == nil {
		t.Fatalf("体重为%f，预期返回体重不能为0的错误，但是未返回", c.Weight)
	}
	c = input.Person{
		Weight: -1,
	}
	err = Calc{}.BMI(&c)
	if err == nil {
		t.Fatalf("体重为%f，预期返回体重不能为负数的错误，但是未返回", c.Weight)
	}
	c = input.Person{
		Weight: 65,
		Tall:   0,
	}
	err = Calc{}.BMI(&c)
	if err == nil {
		t.Fatalf("身高为%f，预期返回身高不能为0的错误，但是未返回", c.Tall)
	}
	c = input.Person{
		Tall: -1,
	}
	err = Calc{}.BMI(&c)
	if err == nil {
		t.Fatalf("身高为%f，预期返回身高不能为负数的错误，但是未返回", c.Tall)
	}
}

func TestCalc_FatRate(t *testing.T) {
	c := input.Person{
		Age: 0,
		Sex: "男",
		Bmi: 20,
	}
	err := Calc{}.FatRate(&c)
	if err == nil {
		t.Fatalf("年龄为%d，预期返回年龄不能为0的错误，但是未返回", c.Age)
	}
	c = input.Person{
		Age: -1,
	}
	err = Calc{}.FatRate(&c)
	if err == nil {
		t.Fatalf("年龄为%d，预期返回年龄不能为负数的错误，但是未返回", c.Age)
	}
	c = input.Person{
		Age: 151,
	}
	err = Calc{}.FatRate(&c)
	if err == nil {
		t.Fatalf("年龄为%d，预期返回年龄不能超过150的错误，但是未返回", c.Age)
	}
	c = input.Person{
		Age: 18,
		Sex: "男",
		Bmi: 0,
	}
	err = Calc{}.FatRate(&c)
	if err == nil {
		t.Fatalf("Bmi为%f，预期返回Bmi不能为0的错误，但是未返回", c.Bmi)
	}
	c = input.Person{
		Bmi: -1,
	}
	err = Calc{}.FatRate(&c)
	if err == nil {
		t.Fatalf("Bmi为%f，预期返回Bmi不能为负数的错误，但是未返回", c.Bmi)
	}
	c = input.Person{
		Age: 18,
		Sex: "性别不明",
		Bmi: 20,
	}
	err = Calc{}.FatRate(&c)
	if err == nil {
		t.Fatalf("Sex为%v，预期返回Sex不能为0的错误，但是未返回", c.Sex)
	}
}
