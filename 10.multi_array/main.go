package main

import "fmt"

func main() {
	frSvc := &fatRatService{s: GetFatRateSuggestion()}
	//fakePerson:=getFakePersonInfo()
	//fmt.Println(frSvc.GiveSuggestionToPerson(fakePerson))
	for {
		p := GetMaterialFromInput()
		fmt.Println(frSvc.GiveSuggestionToPerson(p))
	}
}
