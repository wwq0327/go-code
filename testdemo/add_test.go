package testdemo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdd(t *testing.T) {

	Convey("input 1 equae 1", t, func() {
		So(Add(1), ShouldEqual, 1)
	})
}

func TestString(t *testing.T) {

	Convey("输入数字1，输出字符1", t, func() {
		So(I2S(1), ShouldEqual, "1")
		So(I2S(123), ShouldEqual, "123")
	})
}

func TestDiv(t *testing.T) {
	Convey("如果a＝6， b＝2，则输入3", t, func() {
		v, _ := Div(6, 2)
		So(v, ShouldEqual, 3)

	})

}
