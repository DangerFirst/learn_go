package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {
	r := &FatRateRank{}
	num := 1000
	//录入基本数据
	for i := 0; i < num; i++ {
		rand.Seed(time.Now().UnixNano())
		r.inputRecord(strconv.Itoa(i), float64(rand.Intn(40))/100)
		time.Sleep(10 * time.Millisecond)
	}
	for _, item := range r.items {
		fatRate := <-item.FatRate
		fmt.Printf("name:%s,faterate:%.2f\n", item.Name, fatRate)
		item.FatRate <- fatRate
	}
	//不断获取排名
	go func() {
		for {
			wg := sync.WaitGroup{}
			wg.Add(num)
			for _, item := range r.items {
				go func(f *RandItem) {
					defer wg.Done()
					rank, fatRate := r.getRank(f.Name)
					fmt.Printf("%s的现体脂率为%.2f,排名为%d\n", f.Name, fatRate, rank)
				}(item)
			}
			wg.Wait()
		}
	}()
	//更新数据
	go func() {
		for {
			wg := sync.WaitGroup{}
			wg.Add(num)
			for _, item := range r.items {
				go func(f *RandItem) {
					defer wg.Done()
					rand.Seed(time.Now().UnixNano())
					time.Sleep(20 * time.Millisecond)
					addOrReduce := rand.Int() % 2
					tmpFateRate := <-f.FatRate
					if addOrReduce == 0 {
						f.FatRate <- tmpFateRate + float64(rand.Intn(20))/100
					} else {
						tmpFateRate -= float64(rand.Intn(20)) / 100
						if tmpFateRate <= 0 {
							f.FatRate <- 0
						} else if tmpFateRate >= 0.4 {
							f.FatRate <- 0.4
						} else {
							f.FatRate <- tmpFateRate
						}
					}
					rank, fatRate := r.getRank(f.Name)
					fmt.Printf("%s体脂率调整后，体脂率为%.2f,排名为%d\n", f.Name, fatRate, rank)
				}(item)
			}
			wg.Wait()
		}
	}()
	time.Sleep(1 * time.Minute)
}

type RandItem struct {
	Name    string
	FatRate chan float64
}

type FatRateRank struct {
	items []*RandItem
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
			fatRate := <-item.FatRate
			if fatRate > minFatRate {
				item.FatRate <- minFatRate
			} else {
				item.FatRate <- fatRate
			}
			r.items[i] = item
			found = true
			break
		}
	}
	if !found {
		randItem := &RandItem{
			Name: name,
		}
		randItem.FatRate = make(chan float64, 1)
		randItem.FatRate <- minFatRate
		r.items = append(r.items, randItem)
	}
}

func (r *FatRateRank) getRank(name string) (rank int, fateRate float64) {
	frs := map[float64]struct{}{}
	for _, item := range r.items {
		tmpFateRate := <-item.FatRate
		frs[tmpFateRate] = struct{}{}
		if item.Name == name {
			fateRate = tmpFateRate
		}
		item.FatRate <- tmpFateRate
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
