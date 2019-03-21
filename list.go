package timesortedlist

import "sort"

// TimeSortedList : Has time series items which are sorted
type TimeSortedList interface {
	sort.Interface
	AddItem(unixTime int64, item interface{})
	AddTimeItem(item *TimeItem)
	Filled() bool
	Cap() int
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

// Len: implementation for sort.Interface
func (tsl *timeSortedList) Len() int {
	return len(tsl.dataList)
}

// Less: implementation for sort.Interface
func (tsl *timeSortedList) Less(i, j int) bool {
	return tsl.dataList[i].UnixTime <= tsl.dataList[j].UnixTime
}

// Swap: implementation for sort.Interface
func (tsl *timeSortedList) Swap(i, j int) {
	tmp := tsl.dataList[i]
	tsl.dataList[i] = tsl.dataList[j]
	tsl.dataList[j] = tmp
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
		oldItem := tsl.dataList[0]
		if oldItem.UnixTime > item.UnixTime {
			// if new item is older than oldest item, ignore
			return
		} else {
			// drop oldest item without changing capacity
			for i := 0; i < tsl.capacity; i++ {
				if i < tsl.capacity-1 {
					// shift to front
					tsl.dataList[i] = tsl.dataList[i+1]
				} else {
					// insert to last
					tsl.dataList[i] = *item
				}
			}
			// TODO: should be inserted more wisely
			sort.Sort(tsl)
		}
	} else if len(tsl.dataList) == 0 {
		// if empty just add
		tsl.dataList = append(tsl.dataList, *item)
	} else {
		tsl.dataList = append(tsl.dataList, *item)
		// TODO: should be inserted more wisely
		sort.Sort(tsl)
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
	//TODO: still needs to be implemented
	idx := sort.Search(len(tsl.dataList), func(i int) bool {
		return fromUnixTime <= tsl.dataList[i].UnixTime
	})
	return tsl.dataList[idx:]
}

func (tsl *timeSortedList) GetItemsUntil(untilUnixTime int64) []TimeItem {
	//TODO: still needs to be implemented
	idx := sort.Search(len(tsl.dataList), func(i int) bool {
		return tsl.dataList[i].UnixTime <= untilUnixTime
	})
	return tsl.dataList[:idx]
}

func (tsl *timeSortedList) GetItemsFromUntil(fromUnixTime, untilUnixTime int64) []TimeItem {
	//TODO: still needs to be implemented
	if fromUnixTime >= untilUnixTime {
		return nil
	}
	fIdx := sort.Search(len(tsl.dataList), func(i int) bool {
		return fromUnixTime <= tsl.dataList[i].UnixTime
	})
	uIdx := sort.Search(len(tsl.dataList), func(i int) bool {
		return tsl.dataList[i].UnixTime <= untilUnixTime
	})
	return tsl.dataList[fIdx:uIdx]
}
