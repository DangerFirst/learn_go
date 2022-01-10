package main

import (
	"math"
	"sort"
)

type RandItem struct {
	Name    string
	FatRate float64
}

type FatRateRank struct {
	items []RandItem
}

func (r *FatRateRank) inputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
	}
	for i, item := range r.items {
		if item.Name == name {
			if item.FatRate > minFatRate {
				item.FatRate = minFatRate
			}
			r.items[i] = item
		}
	}
}

func (r *FatRateRank) getRank(name string) (rank int, fateRate float64) {
	fateRate2PersonMap := map[float64][]string{}
	rankArr := make([]float64, 0, len(personFatRate))
	for nameItem, frItem := range personFatRate {
		fateRate2PersonMap[frItem] = append(fateRate2PersonMap[frItem], nameItem)
		rankArr = append(rankArr, frItem)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		_names := fateRate2PersonMap[frItem]
		for _, _name := range _names {
			if _name == name {
				rank = i + 1
				fateRate = frItem
				return
			}
		}
	}
	return
}
