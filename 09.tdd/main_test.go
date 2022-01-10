package main

import "testing"

func TestCase1Part1(t *testing.T) {
	//inputRecord("王强", 0.38)
	inputRecord("王强", 0.32)
	inputRecord("王强", 0.38)
	{
		randOfWQ, fatRateOfWQ := getRank("王强")
		if randOfWQ != 1 {
			t.Fatalf("预期王强第一，但是得到的是%d", randOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂率为0.32，但是得到的是%f", fatRateOfWQ)
		}
	}
}

func TestCase1(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("王强", 0.32)
	inputRecord("李静", 0.28)
	{
		randOfWQ, fatRateOfWQ := getRank("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期王强第二，但是得到的是%d", randOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂率为0.32，但是得到的是%f", fatRateOfWQ)
		}
	}
	{
		randOfLJ, fatRateOfLJ := getRank("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期李静第一，但是得到的是%d", randOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期李静的体脂率为0.28，但是得到的是%f", fatRateOfLJ)
		}
	}
}

func TestCase2(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("张伟", 0.38)
	inputRecord("李静", 0.28)
	{
		randOfWQ, fatRateOfWQ := getRank("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期王强第二，但是得到的是%d", randOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂率为0.32，但是得到的是%f", fatRateOfWQ)
		}
	}
	{
		randOfZW, fatRateOfZW := getRank("张伟")
		if randOfZW != 3 {
			t.Fatalf("预期张伟第三，但是得到的是%d", randOfZW)
		}
		if fatRateOfZW != 0.38 {
			t.Fatalf("预期张伟的体脂率为0.38，但是得到的是%f", fatRateOfZW)
		}
	}
	{
		randOfLJ, fatRateOfLJ := getRank("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期李静第一，但是得到的是%d", randOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期李静的体脂率为0.28，但是得到的是%f", fatRateOfLJ)
		}
	}
}

func TestCase3(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("张伟")
	inputRecord("李静", 0.28)
	{
		randOfWQ, fatRateOfWQ := getRank("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期王强第二，但是得到的是%d", randOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂率为0.32，但是得到的是%f", fatRateOfWQ)
		}
	}
	{
		randOfZW, _ := getRank("张伟")
		if randOfZW != 3 {
			t.Fatalf("预期张伟第三，但是得到的是%d", randOfZW)
		}
	}
	{
		randOfLJ, fatRateOfLJ := getRank("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期李静第一，但是得到的是%d", randOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期李静的体脂率为0.28，但是得到的是%f", fatRateOfLJ)
		}
	}
}
