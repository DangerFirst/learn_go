package main

import "sync"

type rank struct {
	standard []string
}

var globalRank = &rank{}

//var initGlobalRankStandardInit bool = false
//var initGlobalRankStandardLock sync.Mutex
var once sync.Once

//func init() {
//	globalRank.standard = []string{"Asia"}
//}

func initGlobalRankStandard(standard []string) {
	once.Do(func() {
		globalRank.standard = standard
	})
	//initGlobalRankStandardLock.Lock()
	//defer initGlobalRankStandardLock.Unlock()
	//if initGlobalRankStandardInit {
	//	return
	//}
	//initGlobalRankStandardInit = true

}

func main() {
	standard := []string{"asia"}
	for i := 0; i < 10; i++ {
		go func() {
			initGlobalRankStandard(standard)
		}()
	}
}
