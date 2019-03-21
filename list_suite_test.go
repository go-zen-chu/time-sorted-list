package timesortedlist

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTimeSortedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TimeSortedList Suite")
}
