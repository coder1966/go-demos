package collect

import "time"

type Collector interface {
	GetTime() int64
	Stop()
}

func NewCollect() Collector {
	return &Collect{}
}

type Collect struct {
	ts int64
}

func (c *Collect) GetTime() int64 {
	c.ts = time.Now().Unix()
	return c.ts
}

func (c *Collect) Stop() {}
