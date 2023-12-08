package feed

import "fmt"

type Feeder interface {
	Feed(int64) error
	Clear()
}

func NewFeed() Feeder {
	return &Feed{}
}

type Feed struct{}

func (f *Feed) Feed(ts int64) error {
	// f.ts = ts
	fmt.Println("Yes, feeding: ", ts)
	return nil
}

func (f *Feed) Clear() {}
