package main

type Assets struct {
	assets []Asset
}

func (a *Assets) DoStarWork() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok {
			d.Open()
		}
	}
}

func (a *Assets) DoStopWork() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok {
			d.Close()
		}
	}
}

func (a *Assets) Unlock() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok {
			d.Unlock()
		}
	}
}

func (a *Assets) Lock() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok {
			d.Lock()
		}
	}
}
