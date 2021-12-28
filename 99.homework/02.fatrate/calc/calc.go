package calc

import (
	"github.com/armstrongli/go-bmi"
	"learn.go/99.homework/02.fatrate/input"
	"log"
)

type Calc struct {
}

func (Calc) BMI(person *input.Person) error {
	bmi, err := gobmi.BMI(person.Weight, person.Tall)
	if err != nil {
		log.Println("error when calculating bmi:", err)
		return err
	}
	person.Bmi = bmi
	return nil
}

func (Calc) FatRate(person *input.Person) error {
	person.FatRate = gobmi.Fatrate(person.Age, person.Sex, person.Bmi)
	return nil
}
