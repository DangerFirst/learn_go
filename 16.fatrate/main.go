package main

func main() {
	frSvc := &fatRatService{
		s:      GetFatRateSuggestion(),
		input:  &inputFromStd{},
		output: &StdOut{},
	}
	//fakePerson:=getFakePersonInfo()
	//fmt.Println(frSvc.GiveSuggestionToPerson(fakePerson))
	for {
		p := frSvc.input.GetInput()
		frSvc.GiveSuggestionToPerson(&p)
	}
}
