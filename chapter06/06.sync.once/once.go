package main

import "sync"

type rank struct {
	standard []string
}

var globalRank = &rank{}
var initGlobalRankStandardInit bool = false
var initGlobalRankStandardLock sync.Mutex

func init() {
	globalRank.standard = []string{"Asia"}
}

func initGlobalRankStandard(standard []string) {
	initGlobalRankStandardLock.Lock()
	defer initGlobalRankStandardLock.Unlock()
	if initGlobalRankStandardInit {
		return
	}
	initGlobalRankStandardInit = true
	globalRank.standard = standard
}

func main() {
	standard := []string{"asia"}
	for i := 0; i < 10; i++ {
		go func() {
			initGlobalRankStandard(standard)
		}()
	}
}
