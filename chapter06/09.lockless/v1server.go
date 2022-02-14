package main

import (
	"fmt"
	"sync"
	"time"
)

type WebServerV1 struct {
	config     Config
	ConfigLock sync.RWMutex
}

func (ws *WebServerV1) Reload() {
	ws.ConfigLock.Lock()
	defer ws.ConfigLock.Unlock()
	ws.config.Content = fmt.Sprintf("%d", time.Now().UnixNano())
}

func (ws *WebServerV1) ReloadWorker() {
	for {
		time.Sleep(10 * time.Millisecond)
		ws.Reload()
	}
}

func (ws *WebServerV1) Visit() string {
	ws.ConfigLock.RLock()
	defer ws.ConfigLock.RUnlock()
	return ws.config.Content
}
