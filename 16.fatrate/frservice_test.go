package main

import (
	"testing"
)

func TestFrServiceSuggestion(t *testing.T) {
	realOutput := &fakeOutput{}
	frSvc := &fatRatService{
		s:      GetFatRateSuggestion(),
		input:  &inputFromFake{},
		output: realOutput,
	}
	p := frSvc.input.GetInput()
	expectedOutput := &fakeOutput{
		p: p,
		s: "标准",
	}
	frSvc.GiveSuggestionToPerson(&p)
	if realOutput.s != expectedOutput.s {
		t.Fatalf("预期：%s,实际输出：%s", expectedOutput.s, realOutput.s)
	}
}
