/*
Package timesortedlist is a package of a data structure storing time series data

## What this data structure for?

- in memory time series data store
- any data sorted in time series
- query data using from or until in unixtime

## Usage

### Searching data through unixtime

Since this data structure intented to store time series data, it has abilty to search item via unix time.

```
timeItems := tsl.GetItemsFrom(unixTime)
timeItems := tsl.GetItemsUntil(unixTime)
```

### Add item

When you add item, it will automatically sort

```
tsl.AddItem(unixTime, timeItem)
```
*/
package timesortedlist

import (
	"sort"
)

// TimeSortedList holds time series items which are sorted.
// Inserted items can be obtained by specifying unix time.
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

// TimeItem is a struct for storing item with time infomation.
type TimeItem struct {
	UnixTime int64
	Item     interface{}
}

type timeSortedList struct {
	dataList []TimeItem
	capacity int
}

// NewTimeSortedList initializes TimeSortedList.
// capacity is
func NewTimeSortedList(capacity int) TimeSortedList {
	l := make([]TimeItem, 0, capacity)
	return &timeSortedList{
		dataList: l,
		capacity: capacity,
	}
}

// Len gets actual length of list.
func (tsl *timeSortedList) Len() int {
	return len(tsl.dataList)
}

// Cap gets initialized capacity.
// If length of list is same as Cap then the list is filled.
func (tsl *timeSortedList) Cap() int {
	return tsl.capacity
}

// AddItem adds any structure with specified time.
func (tsl *timeSortedList) AddItem(unixTime int64, item interface{}) {
	ti := &TimeItem{
		UnixTime: unixTime,
		Item:     item,
	}
	tsl.AddTimeItem(ti)
}

// AddTimeItem adds TimeItem with time ordered.
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

// Filled checks if the list is filled.
func (tsl *timeSortedList) Filled() bool {
	// by comparing with defined capacity, make the list fix sized
	return len(tsl.dataList) == tsl.capacity
}

// GetItem gets item with specified index.
func (tsl *timeSortedList) GetItem(idx int) *TimeItem {
	if len(tsl.dataList) == 0 || len(tsl.dataList) <= idx {
		// empty or out of range
		return nil
	}
	return &tsl.dataList[idx]
}

// GetItemsFrom gets item from specified time
func (tsl *timeSortedList) GetItemsFrom(fromUnixTime int64) []TimeItem {
	idx := sort.Search(len(tsl.dataList), func(i int) bool {
		return fromUnixTime <= tsl.dataList[i].UnixTime
	})
	return tsl.dataList[idx:]
}

// GetItemsFromUntil gets item until specified time
func (tsl *timeSortedList) GetItemsUntil(untilUnixTime int64) []TimeItem {
	idx := sort.Search(len(tsl.dataList), func(i int) bool {
		// get index where it surpass until time
		return untilUnixTime < tsl.dataList[i].UnixTime
	})
	return tsl.dataList[:idx]
}

// GetItemsFromUntil get items with specified time range
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
