package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrimeFactors(t *testing.T) {
	Convey("Testing PrimeFactors", t, func() {
		So(PrimeFactors(40), ShouldResemble, []int{2, 4, 5})
	})
}
