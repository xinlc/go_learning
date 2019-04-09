package bdd_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey" // . 代码引入到当前空间，例如：就不用加fmt
)

// go get -u github.com/smartystreets/goconvey
// $GOPATH/bin/goconvey 启动web UI
func TestSpec(t *testing.T) {
	Convey("Given 2 even numbers", t, func() {
		a := 2
		b := 4

		Convey("When add the two numbers", func() {
			c := a + b

			Convey("Then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}
