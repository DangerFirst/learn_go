package main

import "sort"

var (
	personFatRate = map[string]float64{}
)

func inputRecord(name string, fatRate float64) {
	personFatRate[name] = fatRate
}

func getRand(name string) (rand int, fateRate float64) {
	fateRate2PersonMap := map[float64]string{}
	randArr := make([]float64, 0, len(personFatRate))
	for nameItem, frItem := range personFatRate {
		fateRate2PersonMap[frItem] = nameItem
		randArr = append(randArr, frItem)
	}
	sort.Float64s(randArr)
	for i, frItem := range randArr {
		_name := fateRate2PersonMap[frItem]
		if _name == name {
			rand = i + 1
			fateRate = frItem
			return
		}
	}
	return
}
