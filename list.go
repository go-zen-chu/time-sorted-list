package timesortedlist

type TimeSortedList interface {
	Length() int
}

type timeSortedList struct {
	dataList []interface{}
}

func NewTimeSortedList(length int) TimeSortedList {
	l := make([]interface{}, length)
	return &timeSortedList{
		dataList: l,
	}
}

func (tsl *timeSortedList) Length() int {
	return len(tsl.dataList)
}
