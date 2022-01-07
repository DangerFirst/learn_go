package main

func GetFatRateSuggestion() *fateRateSuggestion {
	return &fateRateSuggestion{
		suggArr: [][][]int{
			{ //第一个元素，表示男
				{ //18~39
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
				},
				{ //40~59

				},
				{ //60+

				},
			},
			{ //第二个元素，表示女
				{ //18~39

				},
				{ //40~59

				},
				{ //60+

				},
			},
		},
	}
}

type fateRateSuggestion struct {
	suggArr [][][]int
}

func (s *fateRateSuggestion) GetSuggestion(person *Person) string {
	sexIdx := s.getIndexOfSex(person.Sex)
	ageIdx := s.getIndexOfAge(person.Age)
	maxFRSupported := len(s.suggArr[sexIdx][ageIdx]) - 1
	frIdxz := int(person.FatRate * 100)
	if frIdxz >= maxFRSupported {
		frIdxz = maxFRSupported
	}
	suggIdx := s.suggArr[sexIdx][ageIdx][frIdxz]
	return s.translateResult(suggIdx)
}

func (s *fateRateSuggestion) getIndexOfSex(sex string) int {
	if sex == "男" {
		return 0
	}
	return 1
}

func (s *fateRateSuggestion) getIndexOfAge(age int) int {
	switch {
	case age >= 18 && age < 40:
		return 0
	case age >= 40 && age < 60:
		return 1
	case age >= 60:
		return 2
	default:
		return -1

	}
}

func (s *fateRateSuggestion) translateResult(idx int) string {
	switch idx {
	case 0:
		return "偏瘦"
	case 1:
		return "标准"
	case 2:
		return "偏重"
	case 3:
		return "肥胖"
	case 4:
		return "过度肥胖"
	}
	return "未知"
}
