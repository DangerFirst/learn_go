package main

import "testing"

func TestCase1(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("王强", 0.32)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 1 {
			t.Fatalf("预期王强第一，但是得到的是%d", randOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂率为0.32，但是得到的是%f", fatRateOfWQ)
		}
	}
	inputRecord("李静", 0.28)
	{
		randOfLJ, fatRateOfLJ := getRand("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期李静第一，但是得到的是%d", randOfLJ)
		}
		if fatRateOfLJ != 0.32 {
			t.Fatalf("预期李静的体脂率为0.32，但是得到的是%f", fatRateOfLJ)
		}
	}
}

func TestCase2(t *testing.T) {

}
