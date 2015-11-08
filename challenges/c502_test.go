package challenges

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestC502(t *testing.T) {
	Convey("Testing C502", t, func() {
		Convey("height=2", FailureContinues, func() {
			So(c502ComputeSum(2, 2), ShouldEqual, 3)
			So(c502ComputeSum(3, 2), ShouldEqual, 6)
			So(c502ComputeSum(4, 2), ShouldEqual, 10)
			So(c502ComputeSum(5, 2), ShouldEqual, 16)
			So(c502ComputeSum(6, 2), ShouldEqual, 28)
			So(c502ComputeSum(7, 2), ShouldEqual, 56)
			So(c502ComputeSum(8, 2), ShouldEqual, 120)
			So(c502ComputeSum(9, 2), ShouldEqual, 256)
			So(c502ComputeSum(10, 2), ShouldEqual, 528)
		})
		Convey("height=3", FailureContinues, func() {
			So(c502ComputeSum(2, 3), ShouldEqual, 3)
			So(c502ComputeSum(3, 3), ShouldEqual, 9)
			So(c502ComputeSum(4, 3), ShouldEqual, 31)
			/*
				So(c502ComputeSum(5, 3), ShouldEqual, 16)
				So(c502ComputeSum(6, 3), ShouldEqual, 28)
				So(c502ComputeSum(7, 3), ShouldEqual, 56)
				So(c502ComputeSum(8, 3), ShouldEqual, 120)
				So(c502ComputeSum(9, 3), ShouldEqual, 256)
				So(c502ComputeSum(10, 3), ShouldEqual, 528)
			*/
		})
		Convey("height=4", FailureContinues, func() {
			So(c502ComputeSum(2, 4), ShouldEqual, 10)
			// So(c502ComputeSum(3, 4), ShouldEqual, 32)
			// So(c502ComputeSum(4, 4), ShouldEqual, 112)
			/*
				So(c502ComputeSum(5, 4), ShouldEqual, 16)
				So(c502ComputeSum(6, 4), ShouldEqual, 28)
				So(c502ComputeSum(7, 4), ShouldEqual, 56)
				So(c502ComputeSum(8, 4), ShouldEqual, 120)
				So(c502ComputeSum(9, 4), ShouldEqual, 256)
				So(c502ComputeSum(10, 4), ShouldEqual, 528)
			*/
		})
	})
}
