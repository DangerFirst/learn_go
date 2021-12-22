package calc

func CalcuBmi(weight, tall float64) (bmi float64) {
	//计算BMI
	bmi = weight / ((tall * tall) / 10000)
	return
}


