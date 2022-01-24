package calc

import (
	"fatrate/input"
	"github.com/armstrongli/go-bmi"
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
	fatRate, err := gobmi.Fatrate(person.Age, person.Sex, person.Bmi)
	if err != nil {
		log.Println("error when calculating fatRate:", err)
		return err
	}
	person.FatRate = fatRate
	return nil
}
