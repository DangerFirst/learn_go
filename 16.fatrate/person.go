package main

import (
	"github.com/armstrongli/go-bmi"
	"log"
)

type Person struct {
	Name, Sex                  string
	Age                        int
	Tall, Weight, Bmi, FatRate float64
}

func (p *Person) calcBmi() error {
	bmi, err := gobmi.BMI(p.Weight, p.Tall)
	if err != nil {
		log.Printf("error when calculating Bmi for Personp[%s]：%v", p.Name, err)
		return err
	}
	p.Bmi = bmi
	return nil
}

func (p *Person) calcFatRate() error {
	fatRate, err := gobmi.Fatrate(p.Age, p.Sex, p.Bmi)
	if err != nil {
		log.Printf("error when calculating FatRate for Personp[%s]：%v", p.Name, err)
		return err
	}
	p.FatRate = fatRate
	return nil
}
