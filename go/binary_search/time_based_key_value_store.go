package binary_search

import "sort"

type Value struct {
	value     string
	timestamp int
}
type TimeMap struct {
	mp map[string][]Value
}

func Constructor() TimeMap {
	return TimeMap{make(map[string][]Value)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.mp[key]; !ok {
		this.mp[key] = make([]Value, 0)
	}
	this.mp[key] = append(this.mp[key], Value{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	i := sort.Search(len(this.mp[key]), func(j int) bool {
		return this.mp[key][j].timestamp > timestamp
	})

	if i == 0 {
		return ""
	}
	return this.mp[key][i-1].value
}
