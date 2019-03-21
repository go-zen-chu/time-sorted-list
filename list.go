package timesortedlist

import "sort"

type TimeSortedList interface {
	sort.Interface
	AddItem(unixTime int64, item interface{})
	AddTimeItem(item TimeItem)
	Filled() bool
}

type TimeItem struct {
	UnixTime int64
	Item     interface{}
}

type timeSortedList struct {
	dataList []TimeItem
}

// NewTimeSortedList : Initialize TimeSortedList
func NewTimeSortedList(capacity int) TimeSortedList {
	l := make([]TimeItem, 0, capacity)
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

func (tsl *timeSortedList) AddItem(unixTime int64, item interface{}) {
	ti := &TimeItem{
		UnixTime: unixTime,
		Item:     item,
	}
	tsl.AddTimeItem(ti)
}

func (tsl *timeSortedList) AddTimeItem(item *TimeItem) {
	if tsl.Filled() {

	} else if len(tsl.dataList) == 0 {
		// if empty just add
		tsl.dataList = append(tsl.dataList, item)
	} else {
		tsl.dataList = append(tsl.dataList, item)
		// TODO: should be inserted more wisely
		sort.Sort(tsl)
	}
}

func (tsl *timeSortedList) Filled() bool {
	return len(tsl.dataList) == cap(tsl.dataList)
}
