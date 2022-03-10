package main

import (
	"fatrate-rank/pkg/apis"
	"github.com/armstrongli/go-bmi"
	"log"
)

type Calc struct {
}

func (c *Calc) BMI(person *apis.PersonalInformation) (bmi float64, err error) {
	bmi, err = gobmi.BMI(person.Weight, person.Tall)
	if err != nil {
		log.Println("error when calculating bmi:", err)
		return -1, err
	}
	return bmi, nil
}

func (c *Calc) FatRate(person *apis.PersonalInformation) (fatRate float64, err error) {
	bmi, _ := c.BMI(person)
	fatRate, err = gobmi.Fatrate(person.Age, person.Sex, bmi)
	if err != nil {
		log.Println("error when calculating fatRate:", err)
		return -1, err
	}
	return fatRate, nil
}
