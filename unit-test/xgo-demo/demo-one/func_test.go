package demo_one

import (
	"fmt"
	"github.com/xhd2015/xgo/runtime/mock"
	"testing"
)

func TestMyFunc(t *testing.T) {
	mock.Patch(MyFunc, func() string {
		return "Hello, Golang!"
	})
	text := MyFunc()

	fmt.Println(text)
}
