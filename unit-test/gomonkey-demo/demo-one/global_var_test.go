package demo_one

import (
	"github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var num = 10 //全局变量

func TestApplyGlobalVar(t *testing.T) {
	Convey("TestApplyGlobalVar", t, func() {

		// 测试将全局变量 num 修改为 150。
		Convey("change", func() {
			patches := gomonkey.ApplyGlobalVar(&num, 150)
			defer patches.Reset()
			So(num, ShouldEqual, 150)
		})
		// 测试将全局变量 num 恢复到原始值。
		Convey("recover", func() {
			So(num, ShouldEqual, 10)
		})
	})
}
