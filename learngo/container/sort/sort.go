package main

import (
	"fmt"
	"sort"
	"time"
)

type base struct {
	weight   int
	createAt time.Time
}

type bases []base

func (b bases) Len() int {
	return len(b)
}

//Less weight desc time unix asc
func (b bases) Less(i, j int) bool {
	if b[i].weight == b[j].weight {
		return b[i].createAt.Unix() < b[j].createAt.Unix()
	}
	return b[i].weight > b[j].weight
}

func (b bases) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

//进一步封装
type baseWrapper struct {
	bases  []base
	option func(p, q *base) bool
}

func (b baseWrapper) Len() int {
	return len(b.bases)
}

func (b baseWrapper) Less(i, j int) bool {
	return b.option(&b.bases[i], &b.bases[j])
}

func (b baseWrapper) Swap(i, j int) {
	b.bases[i], b.bases[j] = b.bases[j], b.bases[i]
}

type sortBy func(p, q *base) bool

func sortBaseWrapper(bases []base, by sortBy) {
	sort.Sort(baseWrapper{bases: bases, option: by})
}

func main() {
	sl := []base{
		{10, time.Now()},
		{11, time.Unix(time.Now().Unix()-10000, 0)},
		{12, time.Unix(time.Now().Unix()-10000, 0)},
		{11, time.Unix(time.Now().Unix()-1000, 0)},
	}
	sortBaseWrapper(sl, func(p, q *base) bool {
		if p.weight == q.weight {
			return p.createAt.Unix() < q.createAt.Unix()
		}
		return p.weight > q.weight
	})

	//替代方法
	slc := []base{
		{10, time.Now()},
		{11, time.Unix(time.Now().Unix()-10000, 0)},
		{12, time.Unix(time.Now().Unix()-10000, 0)},
		{11, time.Unix(time.Now().Unix()-1000, 0)},
	}
	sort.Slice(slc, func(i, j int) bool {
		return slc[i].createAt.After(slc[j].createAt)
	})
	fmt.Println(sl)
	fmt.Println(slc)
}
