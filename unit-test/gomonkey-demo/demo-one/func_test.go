package demo_one

import (
	"encoding/json"
	"github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Sum(a, b int) (int, error) {
	// do something in remote computer
	c := a + b
	return c, nil
}

func Compute(a, b int) (int, error) {
	sum, err := Sum(a, b)
	return sum, err
}

func ComputeTow(a, b, c int) (int, error) {
	sum, err := Sum(a, b)
	if err != nil {
		// do something
		return 0, err
	}

	rs, err := Compute(c, sum)
	if err != nil {
		return 0, err
	}
	return rs, nil
}

func TestCompute(t *testing.T) {
	patches := gomonkey.ApplyFunc(Sum, func(a, b int) (int, error) {
		return 2, nil
	})

	defer patches.Reset()
	sum, err := Compute(1, 1)
	assert.NoError(t, err)
	assert.Equal(t, 2, sum)
}

func TestComputeConvey(t *testing.T) {
	Convey("one func for success", t, func() {
		patches := gomonkey.ApplyFunc(Sum, func(a, b int) (int, error) {
			return 2, nil
		})

		defer patches.Reset()
		sum, err := Compute(1, 1)
		So(err, ShouldBeNil)
		So(sum, ShouldEqual, 2)
	})

	Convey("one func for fail", t, func() {
		patches := gomonkey.ApplyFunc(Sum, func(a, b int) (int, error) {
			return 2, nil
		})

		defer patches.Reset()
		sum, err := Compute(1, 2)
		So(err, ShouldBeNil)
		So(sum, ShouldEqual, 3)
	})

	Convey("two func for success", t, func() {
		patches := gomonkey.ApplyFunc(Sum, func(a, b int) (int, error) {
			return 6, nil
		})
		defer patches.Reset()
		patches.ApplyFunc(ComputeTow, func(a, b, c int) (int, error) {
			return 4, nil
		})

		sum, err := ComputeTow(1, 1, 2)
		So(err, ShouldBeNil)
		So(sum, ShouldEqual, 4)

	})

	Convey("input and output param", t, func() {
		patches := gomonkey.ApplyFunc(json.Unmarshal, func(data []byte, v interface{}) error {
			if data == nil {
				panic("input param is nil!")
			}
			p := v.(*map[int]int)
			*p = make(map[int]int)
			(*p)[1] = 2
			(*p)[2] = 4
			return nil
		})
		defer patches.Reset()
		var m map[int]int
		err := json.Unmarshal([]byte("123"), &m)

		So(err, ShouldEqual, nil)
		So(m[1], ShouldEqual, 2)

	})
}
