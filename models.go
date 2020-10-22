package sysmon

import (
	"sync"
	"time"
)

type WatchInfo struct {
	Name  string  `json:"type"`
	Ts    int64   `json:"ts"`
	Value float64 `json:"value"`
}

type Watcher interface {
	Get() ([]WatchInfo, error)
}

type Core struct {
	sync.RWMutex
	series    map[string][]WatchInfo
	timeDelta int64
}

func NewCore(timeDelta int64) *Core {
	s := make(map[string][]WatchInfo)
	return &Core{timeDelta: timeDelta, series: s}
}

func (c *Core) Append(w Watcher) error {
	t := time.Tick(time.Second*1)
	for {
		_ = <-t
		arr,err := w.Get()
		if err != nil {
			return err
		}
		for _, el := range arr{
			c.Add(el)
		}
	}
}

func (c *Core) Add(wi WatchInfo) {
	c.Lock()
	defer c.Unlock()
	tsn := time.Now().Unix()
	var cropFrom int
	for i, el := range c.series[wi.Name] {
		if (tsn-el.Ts) < c.timeDelta {
			cropFrom = i
			break
		}
	}
	if cropFrom > 0 {
		c.series[wi.Name] = c.series[wi.Name][cropFrom:]
	}
	c.series[wi.Name] = append(c.series[wi.Name], wi)
}

func (c *Core) Avg(key string) float64 {
	var res float64
	c.RLock()
	defer c.RUnlock()
	tsn := time.Now().Unix()
	if v, ok := c.series[key]; ok {
		cntr := 0
		for _, el := range v {
			if tsn-el.Ts < c.timeDelta {
				res += el.Value
				cntr++
			}
		}
		if cntr != 0{
			res = res / float64(cntr)
		}
	}
	return res
}
