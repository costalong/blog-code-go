package mock

import (
	demo_one "github.com/costalong/blog-code-go/unit-test/gomock-demo/demo-one"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetUserById(t *testing.T) {
	//初始化一个mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := NewMockUserSrv(ctrl)
	// mock方法调用次数支持如下
	mockObj.EXPECT().GetUserById(gomock.Any()).Return("11").AnyTimes()
	str := mockObj.GetUserById(11)

	// 判断结果
	assert.NotEmpty(t, str)    // 断言 str 不为空
	assert.Equal(t, "11", str) // 断言 str 的值等于 ”11“
}

// 测试 update

func Test_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := NewMockUserSrv(ctrl)

	mockObj.EXPECT().GetUserById(gomock.Any()).Return("11").AnyTimes()

	//mockObj.EXPECT().Update(gomock.Any(), gomock.Any()).Return("1", nil).AnyTimes()

	src := demo_one.NewUserService()
	str, err := src.Update(11, "11")
	assert.NoError(t, err)
	assert.NotEmpty(t, str)
	assert.Equal(t, "11", str)
}
