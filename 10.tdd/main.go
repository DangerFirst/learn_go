package main

import "sort"

var (
	personFatRate = map[string]float64{}
)

func inputRecord(name string, fatRate float64) {
	personFatRate[name] = fatRate
}

func getRank(name string) (rank int, fateRate float64) {
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
