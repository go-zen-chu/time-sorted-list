package timesortedlist_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/go-zen-chu/time-sorted-list"
)

type SampleStruct struct {
	Sample string
}

var _ = Describe("TimeSortedList", func() {
	nowUnixTime := time.Now().Unix()
	sampleTimeItems := []TimeItem{
		TimeItem{
			UnixTime: nowUnixTime,
			Item:     SampleStruct{Sample: "sample"},
		},
		TimeItem{
			UnixTime: nowUnixTime + 1,
			Item:     SampleStruct{Sample: "sample1"},
		},
		TimeItem{
			UnixTime: nowUnixTime + 2,
			Item:     SampleStruct{Sample: "sample2"},
		},
		TimeItem{
			UnixTime: nowUnixTime + 3,
			Item:     SampleStruct{Sample: "sample3"},
		},
		TimeItem{
			UnixTime: nowUnixTime + 4,
			Item:     SampleStruct{Sample: "sample4"},
		},
	}

	Describe("NewTimeSortedList", func() {
		It("should create len 0 cap 5 list", func() {
			tsl := NewTimeSortedList(5)
			Expect(tsl.Len()).To(Equal(0))
			Expect(tsl.Cap()).To(Equal(5))
		})
	})

	Describe("Filled", func() {
		It("should be not filled in first place", func() {
			tsl := NewTimeSortedList(5)
			Expect(tsl.Filled()).To(BeFalse())
		})

		It("should be filled", func() {
			tsl := NewTimeSortedList(5)
			for _, ti := range sampleTimeItems {
				tsl.AddTimeItem(&ti)
			}
			Expect(tsl.Filled()).To(BeTrue())
		})
	})
})
