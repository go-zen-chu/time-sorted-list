package timesortedlist

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type sampleStruct struct {
	sample string
}

var (
	nowUnixTime     = time.Now().Unix()
	sampleTimeItems = []TimeItem{
		TimeItem{
			UnixTime: nowUnixTime,
			Item:     sampleStruct{sample: "sample"},
		},
		TimeItem{
			UnixTime: nowUnixTime + 1,
			Item:     sampleStruct{sample: "sample1"},
		},
		TimeItem{
			UnixTime: nowUnixTime + 2,
			Item:     sampleStruct{sample: "sample2"},
		},
		TimeItem{
			UnixTime: nowUnixTime + 3,
			Item:     sampleStruct{sample: "sample3"},
		},
		TimeItem{
			UnixTime: nowUnixTime + 4,
			Item:     sampleStruct{sample: "sample4"},
		},
	}
)

func genSampleTimeSortedList() TimeSortedList {
	tsl := NewTimeSortedList(5)
	for _, ti := range sampleTimeItems {
		tsl.AddTimeItem(&ti)
	}
	return tsl
}

var _ = Describe("TimeSortedList", func() {
	Describe("NewTimeSortedList", func() {
		It("should create len 0 cap 5 list", func() {
			tsl := NewTimeSortedList(5)
			Expect(tsl.Len()).To(Equal(0))
			Expect(tsl.Cap()).To(Equal(5))
		})
	})

	Describe("AddItem", func() {
		It("should add item properly", func() {
			tsl := NewTimeSortedList(5)
			tsl.AddItem(nowUnixTime, sampleTimeItems[0])
			tsl.AddItem(nowUnixTime+1, sampleTimeItems[1])
			Expect(tsl.Len()).To(Equal(2))
			Expect(tsl.Cap()).To(Equal(5))
		})
	})

	Describe("AddTimeItem", func() {
		It("should add item in time order", func() {
			tsl := NewTimeSortedList(5)
			tsl.AddTimeItem(&sampleTimeItems[3])
			tsl.AddTimeItem(&sampleTimeItems[1])
			tsl.AddTimeItem(&sampleTimeItems[4])
			tsl.AddTimeItem(&sampleTimeItems[0])
			Expect(tsl.GetItem(0).Item.(sampleStruct).sample).To(Equal("sample"))
			Expect(tsl.GetItem(2).Item.(sampleStruct).sample).To(Equal("sample3"))
		})

		It("should remove old item if new one added", func() {
			tsl := genSampleTimeSortedList()
			tsl.AddTimeItem(&TimeItem{
				UnixTime: nowUnixTime + 10,
				Item: sampleStruct{
					sample: "sample10",
				},
			})
			Expect(tsl.GetItem(0).Item.(sampleStruct).sample).To(Equal("sample1"))
			Expect(tsl.GetItem(1).Item.(sampleStruct).sample).To(Equal("sample2"))
			Expect(tsl.GetItem(4).Item.(sampleStruct).sample).To(Equal("sample10"))
		})

		It("should ignore if old item is added", func() {
			tsl := genSampleTimeSortedList()
			tsl.AddTimeItem(&TimeItem{
				UnixTime: nowUnixTime - 1,
				Item: sampleStruct{
					sample: "sample-1",
				},
			})
			Expect(tsl.GetItem(0).Item.(sampleStruct).sample).To(Equal("sample"))
		})
	})

	Describe("Filled", func() {
		It("should be not filled in first place", func() {
			tsl := NewTimeSortedList(3)
			Expect(tsl.Filled()).To(BeFalse())
		})

		It("should be filled", func() {
			tsl := genSampleTimeSortedList()
			Expect(tsl.Filled()).To(BeTrue())
		})
	})

	Describe("GetItem", func() {
		It("should get item in range", func() {
			tsl := genSampleTimeSortedList()
			ti := tsl.GetItem(0)
			Expect(ti.Item.(sampleStruct).sample).To(Equal("sample"))
		})

		It("should get nil if TimeSortedList is empty", func() {
			tsl := NewTimeSortedList(3)
			ti := tsl.GetItem(0)
			Expect(ti).To(BeNil())
		})

		It("should get index is out of range", func() {
			tsl := genSampleTimeSortedList()
			ti := tsl.GetItem(100)
			Expect(ti).To(BeNil())
		})
	})
})
