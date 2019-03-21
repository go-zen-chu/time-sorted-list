package timesortedlist

import "sort"

type TimeSortedList interface {
	sort.Interface
}

type TimeItem struct {
	UnixTime int64
	Item     interface{}
}

type timeSortedList struct {
	dataList []TimeItem
}

func NewTimeSortedList(length int) TimeSortedList {
	l := make([]TimeItem, length)
	return &timeSortedList{
		dataList: l,
	}
}

// Len(): implementation for sort.Interface
func (tsl *timeSortedList) Len() int {
	return len(tsl.dataList)
}

func (tsl *timeSortedList) Less(i, j int) bool {
	return tsl.dataList[i].UnixTime <= tsl.dataList[j].UnixTime
}

func (tsl *timeSortedList) Swap(i, j int) {
	tmp := tsl.dataList[i]
	tsl.dataList[i] = tsl.dataList[j]
	tsl.dataList[j] = tmp
}

func (tsl *timeSortedList) Sort() {
	sort.Sort(tsl)
}
