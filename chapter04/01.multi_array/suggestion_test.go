package main

import "testing"

func Test_fateRateSuggestion_GetSuggestion(t *testing.T) {
	sugg := GetFatRateSuggestion()
	tests := []Person{
		{
			Sex:     "男",
			Age:     35,
			FatRate: 0.24,
		},
	}
	if got := sugg.GetSuggestion(&tests[0]); got != "肥胖" {
		t.Fail()
	}
}

func Test_fateRateSuggestion_GetSuggestion1(t *testing.T) {
	type args struct {
		person *Person
	}
	var (
		tests = []struct {
			name string
			args args
			want string
		}{
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.00}}, want: "偏瘦"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.01}}, want: "偏瘦"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.02}}, want: "偏瘦"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.03}}, want: "偏瘦"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.04}}, want: "偏瘦"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.05}}, want: "偏瘦"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.06}}, want: "偏瘦"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.07}}, want: "偏瘦"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.08}}, want: "偏瘦"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.09}}, want: "偏瘦"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.10}}, want: "标准"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.11}}, want: "标准"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.12}}, want: "标准"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.13}}, want: "标准"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.14}}, want: "标准"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.15}}, want: "标准"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.16}}, want: "偏重"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.17}}, want: "偏重"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.18}}, want: "偏重"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.19}}, want: "偏重"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.20}}, want: "偏重"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.21}}, want: "肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.22}}, want: "肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.23}}, want: "肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.24}}, want: "肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.25}}, want: "肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.26}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.27}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.28}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.29}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.30}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.31}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.32}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.33}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.34}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.35}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.36}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.37}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.38}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.39}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.40}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.41}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.42}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.43}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.44}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.45}}, want: "过度肥胖"},
			{name: "0.24-35", args: args{person: &Person{Sex: "男", Age: 35, FatRate: 0.46}}, want: "过度肥胖"},
		}
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := GetFatRateSuggestion()
			if got := s.GetSuggestion(tt.args.person); got != tt.want {
				t.Errorf("GetSuggestion() = %v, want %v", got, tt.want)
			}
		})
	}
}
