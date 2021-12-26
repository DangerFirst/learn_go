package calc

import "testing"

func TestCalcuBmi(t *testing.T) {
	inputWeight, inputTall := 1.0, 100.0
	expectedOutput := 1.0
	t.Logf("开始计算，输入：weight:%f,height:%f,期望结果：%f", inputWeight, inputTall, expectedOutput)
	actualOutput := CalcuBmi(inputWeight, inputTall)
	t.Logf("实际得到结果：%f", actualOutput)
	if actualOutput != expectedOutput {
		t.Errorf("expecting %f,but got %f", expectedOutput, actualOutput)
	}
}
