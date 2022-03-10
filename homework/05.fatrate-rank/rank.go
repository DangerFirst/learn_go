package main

import (
	"math"
	"sort"
)

type RankItem struct {
	Name    string
	FatRate float64
}

type FatRateRank struct {
	items []RankItem
}

func (r *FatRateRank) inputRecord(name string, fatRate ...float64) {
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
		r.items = append(r.items, RankItem{
			Name:    name,
			FatRate: minFatRate,
		})
	}
}

func (r *FatRateRank) getRankWithBubbleSort(name string) (rank int, fateRate float64) {
	for i := 0; i < len(r.items)-1; i++ {
		for j := i + 1; j < len(r.items); j++ {
			if (r.items)[i].FatRate > (r.items)[j].FatRate {
				(r.items)[i], (r.items)[j] = (r.items)[j], (r.items)[i]
			}
		}
	}
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

func (r *FatRateRank) getRankWithQuickSort(name string) (rank int, fateRate float64) {
	quickSort(r.items, 0, len(r.items)-1)
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

func quickSort(items []RankItem, start, end int) {
	pivotIdx := (start + end) / 2
	pivotV := (items)[pivotIdx].FatRate
	l, r := start, end
	for l <= r {
		for (items)[l].FatRate < pivotV {
			l++
		}
		for (items)[r].FatRate > pivotV {
			r--
		}
		if l >= r {
			break
		}
		(items)[l], (items)[r] = (items)[r], (items)[l]
		l++
		r--
	}
	if l == r {
		l++
		r--
	}
	if r > start {
		quickSort(items, start, r)
	}
	if l < end {
		quickSort(items, l, end)
	}
}
