package demo_one

import (
	"context"
	"fmt"
	"github.com/xhd2015/xgo/runtime/core"
	"github.com/xhd2015/xgo/runtime/mock"
	"testing"
)

func TestMethodMock(t *testing.T) {
	myStruct := &MyStruct{name: "my struct"}

	otherStruct := &MyStruct{
		name: "other struct",
	}
	mock.Mock(myStruct.Name, func(ctx context.Context, fn *core.FuncInfo, args core.Object, results core.Object) error {
		results.GetFieldIndex(0).Set("mock struct")
		return nil
	})

	name := myStruct.Name()
	if name != "mock struct" {
		t.Fatalf("expect myStruct.Name() to be 'mock struct', actual: %s", name)
	}
	fmt.Println(name)

	// otherStruct is not affected
	otherName := otherStruct.Name()
	if otherName != "other struct" {
		t.Fatalf("expect otherStruct.Name() to be 'other struct', actual: %s", otherName)
	}
}
