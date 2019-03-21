package timesortedlist_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/go-zen-chu/time-sorted-list"
)

var _ = Describe("TimeSortedList", func() {

	Describe("NewTimeSortedList", func() {
		It("should create len 10 list", func() {
			tsl := NewTimeSortedList(10)
			Expect(tsl.Length()).To(Equal(10))
		})
	})
})
