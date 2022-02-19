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
		fmt.Printf("name:%s,faterate:%.2f\n", item.Name, item.FatRate)
	}
	//不断获取排名
	go func() {
		for {
			wg := sync.WaitGroup{}
			wg.Add(num)
			for _, item := range r.items {
				go func(f *RandItem) {
					f.Lock.RLock()
					defer wg.Done()
					defer f.Lock.RUnlock()
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
					f.Lock.Lock()
					defer wg.Done()
					defer f.Lock.Unlock()
					rand.Seed(time.Now().UnixNano())
					time.Sleep(20 * time.Millisecond)
					addOrReduce := rand.Int() % 2
					if addOrReduce == 0 {
						f.FatRate += float64(rand.Intn(20)) / 100
					} else {
						f.FatRate -= float64(rand.Intn(20)) / 100
						if f.FatRate <= 0 {
							f.FatRate = 0
						} else if f.FatRate >= 0.4 {
							f.FatRate = 0.4
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
	FatRate float64
	Lock    sync.RWMutex
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
			if item.FatRate > minFatRate {
				item.FatRate = minFatRate
			}
			r.items[i] = item
			found = true
			break
		}
	}
	if !found {
		r.items = append(r.items, &RandItem{
			Name:    name,
			FatRate: minFatRate,
		})
	}
}

func (r *FatRateRank) getRank(name string) (rank int, fateRate float64) {
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
