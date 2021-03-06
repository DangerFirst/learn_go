package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"learn.go/chapter11/02.practice/frinterface"
	"learn.go/pkg/apis"
	"log"
	"math"
	"sort"
	"sync"
)

var _ frinterface.ServerInterface = &FatRateRank{}

type RandItem struct {
	Name    string
	Sex     string
	FatRate float64
}

type FatRateRank struct {
	items     []RandItem
	itemsLock sync.Mutex
}

func (r *FatRateRank) RegisterPersonalInformation(pi *apis.PersonalInformation) error {
	bmi, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
	if err != nil {
		log.Printf("计算%s的bmi失败", pi.Name)
		return err
	}
	fr, err := gobmi.Fatrate(int(pi.Age), pi.Sex, bmi)
	if err != nil {
		log.Printf("计算%s的体脂率失败", pi.Name)
		return err
	}
	r.inputRecord(pi.Name, pi.Sex, fr)
	return nil
}

func (r *FatRateRank) UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error) {
	bmi, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
	if err != nil {
		log.Printf("计算%s的bmi失败", pi.Name)
		return nil, err
	}
	fr, err := gobmi.Fatrate(int(pi.Age), pi.Sex, bmi)
	if err != nil {
		log.Printf("计算%s的体脂率失败", pi.Name)
		return nil, err
	}
	r.inputRecord(pi.Name, pi.Sex, fr)
	return &apis.PersonalInformationFatRate{
		Name:    pi.Name,
		FatRate: float32(fr),
	}, nil
}

func (r *FatRateRank) GetFatRate(name string) (*apis.PersonalRank, error) {
	rankId, sex, fr := r.getRank(name)
	return &apis.PersonalRank{
		Name:    name,
		Sex:     sex,
		RankNum: int64(rankId),
		FatRate: float32(fr),
	}, nil
}

func (r *FatRateRank) GetTop() ([]*apis.PersonalRank, error) {
	return r.getRankTop(), nil
}

func NewFatRateRank() *FatRateRank {
	return &FatRateRank{
		items: make([]RandItem, 0, 100),
	}
}

func (r *FatRateRank) inputRecord(name, sex string, fatRate ...float64) {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
	}
	found := false
	for i, item := range r.items {
		if item.Name == name {
			if item.FatRate > minFatRate {
				item.FatRate = minFatRate
			}
			r.items[i] = item
			found = true
			break
		}
	}
	if !found {
		r.items = append(r.items, RandItem{
			Name:    name,
			Sex:     sex,
			FatRate: minFatRate,
		})
	}
}

func (r *FatRateRank) getRank(name string) (rank int, sex string, fateRate float64) {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()
	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate
	})
	frs := map[float64]struct{}{}
	for _, item := range r.items {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fateRate = item.FatRate
		}
	}
	rankArr := make([]float64, 0, len(frs))
	for k := range frs {
		rankArr = append(rankArr, k)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		if frItem == fateRate {
			rank = i + 1
			break
		}
	}
	return
}

func (r *FatRateRank) getRankTop() []*apis.PersonalRank {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()
	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate
	})
	out := make([]*apis.PersonalRank, 0, 10)
	for i := 0; i < 10 && i < len(r.items); i++ {
		item := r.items[i]
		out = append(out, &apis.PersonalRank{
			Name:    item.Name,
			Sex:     item.Sex,
			RankNum: int64(i + 1),
			FatRate: float32(item.FatRate),
		})
	}
	return out
}
