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

func genSampleTimeSortedList() ITimeSortedList {
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

		It("should insert items without changing capacity", func() {
			tsl := genSampleTimeSortedList()
			tsl.AddTimeItem(&TimeItem{
				UnixTime: nowUnixTime + 3,
				Item: sampleStruct{
					sample: "sample3 new",
				},
			})
			tsl.AddTimeItem(&TimeItem{
				UnixTime: nowUnixTime + 1,
				Item: sampleStruct{
					sample: "sample1 new",
				},
			})
			Expect(tsl.GetItem(0).Item.(sampleStruct).sample).To(Equal("sample1 new"))
			Expect(tsl.GetItem(1).Item.(sampleStruct).sample).To(Equal("sample2"))
			Expect(tsl.GetItem(3).Item.(sampleStruct).sample).To(Equal("sample3 new"))
			Expect(tsl.GetItem(4).Item.(sampleStruct).sample).To(Equal("sample4"))
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
		tsl := genSampleTimeSortedList()

		It("should get item in range", func() {
			ti := tsl.GetItem(0)
			Expect(ti.Item.(sampleStruct).sample).To(Equal("sample"))
		})

		It("should get nil if TimeSortedList is empty", func() {
			emptyTsl := NewTimeSortedList(3)
			ti := emptyTsl.GetItem(0)
			Expect(ti).To(BeNil())
		})

		It("should get index is out of range", func() {
			ti := tsl.GetItem(100)
			Expect(ti).To(BeNil())
		})
	})

	Describe("GetItemsFrom", func() {
		tsl := genSampleTimeSortedList()
		It("should get item from specified time", func() {
			tis := tsl.GetItemsFrom(nowUnixTime + 2)
			Expect(len(tis)).To(Equal(3))
			Expect(tis[0].Item.(sampleStruct).sample).To(Equal("sample2"))
			Expect(tis[1].Item.(sampleStruct).sample).To(Equal("sample3"))
		})

		It("should get all items if from time is older than any items", func() {
			tis := tsl.GetItemsFrom(nowUnixTime - 100)
			Expect(len(tis)).To(Equal(5))
			Expect(tis[0].Item.(sampleStruct).sample).To(Equal("sample"))
		})

		It("should get empty array if from time is newer than any items", func() {
			tis := tsl.GetItemsFrom(nowUnixTime + 100)
			Expect(len(tis)).To(Equal(0))
		})
	})

	Describe("GetItemsUntil", func() {
		tsl := genSampleTimeSortedList()
		It("should get item until specified time", func() {
			tis := tsl.GetItemsUntil(nowUnixTime + 2)
			Expect(len(tis)).To(Equal(3))
			Expect(tis[0].Item.(sampleStruct).sample).To(Equal("sample"))
			Expect(tis[1].Item.(sampleStruct).sample).To(Equal("sample1"))
		})

		It("should get empty array if until time is older than any items", func() {
			tis := tsl.GetItemsUntil(nowUnixTime - 100)
			Expect(len(tis)).To(Equal(0))
		})

		It("should get all items if until time is newer than any items", func() {
			tis := tsl.GetItemsUntil(nowUnixTime + 100)
			Expect(len(tis)).To(Equal(5))
			Expect(tis[0].Item.(sampleStruct).sample).To(Equal("sample"))
		})
	})

	Describe("GetItemsFromUntil", func() {
		tsl := genSampleTimeSortedList()

		It("should get empty array if from time is newer than until time", func() {
			tis := tsl.GetItemsFromUntil(nowUnixTime+2, nowUnixTime)
			Expect(len(tis)).To(Equal(0))
		})

		It("should get items from and until specified time", func() {
			tis := tsl.GetItemsFromUntil(nowUnixTime+1, nowUnixTime+3)
			Expect(len(tis)).To(Equal(3))
			Expect(tis[0].Item.(sampleStruct).sample).To(Equal("sample1"))
			Expect(tis[2].Item.(sampleStruct).sample).To(Equal("sample3"))
		})

		It("should get items array when until time is newer than any items", func() {
			tis := tsl.GetItemsFromUntil(nowUnixTime+3, nowUnixTime+100)
			Expect(len(tis)).To(Equal(2))
			Expect(tis[1].Item.(sampleStruct).sample).To(Equal("sample4"))
		})
	})
})
