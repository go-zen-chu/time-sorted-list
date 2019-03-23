package timesortedlist

import (
	"sort"
)

// TimeSortedList : Has time series items which are sorted
type TimeSortedList interface {
	Len() int
	Cap() int
	AddItem(unixTime int64, item interface{})
	AddTimeItem(item *TimeItem)
	Filled() bool
	GetItem(idx int) *TimeItem
	GetItemsFrom(fromUnixTime int64) []TimeItem
	GetItemsUntil(untilUnixTime int64) []TimeItem
	GetItemsFromUntil(fromUnixTime, untilUnixTime int64) []TimeItem
}

// TimeItem : Item with unix time. Any structures are allowed for Item.
type TimeItem struct {
	UnixTime int64
	Item     interface{}
}

type timeSortedList struct {
	dataList []TimeItem
	capacity int
}

// NewTimeSortedList : Initialize TimeSortedList
func NewTimeSortedList(capacity int) TimeSortedList {
	l := make([]TimeItem, 0, capacity)
	return &timeSortedList{
		dataList: l,
		capacity: capacity,
	}
}

func (tsl *timeSortedList) Len() int {
	return len(tsl.dataList)
}

func (tsl *timeSortedList) Cap() int {
	return tsl.capacity
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
		oldestItem := tsl.dataList[0]
		if item.UnixTime < oldestItem.UnixTime {
			// if new item is older than oldest item, ignore
			return
		} else {
			// drop oldest item without changing capacity
			idx := sort.Search(len(tsl.dataList), func(i int) bool {
				return item.UnixTime < tsl.dataList[i].UnixTime
			})
			for i := 0; i < tsl.capacity; i++ {
				if i == idx {
					tsl.dataList[idx-1] = *item
					break // finished inserting. no need to shift
				} else if i == tsl.capacity-1 {
					// if it comes to last, just add item
					tsl.dataList[i] = *item
				} else {
					tsl.dataList[i] = tsl.dataList[i+1]
				}
			}
		}
	} else if len(tsl.dataList) == 0 {
		// if empty just add
		tsl.dataList = append(tsl.dataList, *item)
	} else {
		// insert **after** same unix time item
		idx := sort.Search(len(tsl.dataList), func(i int) bool {
			return item.UnixTime < tsl.dataList[i].UnixTime
		})
		tsl.dataList = append(tsl.dataList[:idx+1], tsl.dataList[idx:]...)
		tsl.dataList[idx] = *item
	}
}

// Filled : Check if the list is filled.
// By comparing with defined capacity, make the list fix sized
func (tsl *timeSortedList) Filled() bool {
	return len(tsl.dataList) == tsl.capacity
}

func (tsl *timeSortedList) GetItem(idx int) *TimeItem {
	if len(tsl.dataList) == 0 || len(tsl.dataList) <= idx {
		// empty or out of range
		return nil
	}
	return &tsl.dataList[idx]
}

func (tsl *timeSortedList) GetItemsFrom(fromUnixTime int64) []TimeItem {
	idx := sort.Search(len(tsl.dataList), func(i int) bool {
		return fromUnixTime <= tsl.dataList[i].UnixTime
	})
	return tsl.dataList[idx:]
}

func (tsl *timeSortedList) GetItemsUntil(untilUnixTime int64) []TimeItem {
	idx := sort.Search(len(tsl.dataList), func(i int) bool {
		// get index where it surpass until time
		return untilUnixTime < tsl.dataList[i].UnixTime
	})
	return tsl.dataList[:idx]
}

func (tsl *timeSortedList) GetItemsFromUntil(fromUnixTime, untilUnixTime int64) []TimeItem {
	if fromUnixTime >= untilUnixTime {
		return make([]TimeItem, 0)
	}
	fIdx := sort.Search(len(tsl.dataList), func(i int) bool {
		return fromUnixTime <= tsl.dataList[i].UnixTime
	})
	uIdx := sort.Search(len(tsl.dataList), func(i int) bool {
		return untilUnixTime < tsl.dataList[i].UnixTime
	})
	return tsl.dataList[fIdx:uIdx]
}
