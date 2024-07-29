package demo_one

import (
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/agiledragon/gomonkey/v2/test/fake"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUserById(id int) string {
	var res string
	if id == 1 {
		res = "1"
	} else {
		res = "2"
	}
	return res
}

func (s *UserService) Update(id int, data interface{}) (string, error) {
	str := s.GetUserById(id)
	return str, nil
}

type Etcd struct {
}

func (e *Etcd) Retrieve(url string) (string, error) {
	output := fmt.Sprintf("%s, %s!", "Hello", "Etcd")
	return output, nil
}

// go test -gcflags="all=-l -N" -v -run TestApplyMethod_One .
func TestApplyMethod_One(t *testing.T) {
	userSrv := NewUserService()

	patches := gomonkey.ApplyMethodReturn(userSrv, "GetUserById", "mock")
	defer patches.Reset()
	str := userSrv.GetUserById(222)
	fmt.Println(str)

}

// go test -gcflags="all=-l -N" -v -run TestApplyMethod_Two .
func TestApplyMethod_Two(t *testing.T) {
	userSrv := NewUserService()
	Convey("TestApplyMethodReturn", t, func() {
		patches := gomonkey.ApplyMethodReturn(userSrv, "GetUserById", "mock")
		defer patches.Reset()
		str := userSrv.GetUserById(1)
		So(str, ShouldEqual, "mock")
	})
}

// go test -gcflags="all=-l -N" -v -run TestApplyMethodReturn .
func TestApplyMethodReturn(t *testing.T) {
	e := &fake.Etcd{}
	Convey("TestApplyMethodReturn", t, func() {
		Convey("declares the values to be returned", func() {
			info1 := "hello cpp"
			patches := gomonkey.ApplyMethodReturn(e, "Retrieve", info1, nil)
			defer patches.Reset()
			for i := 0; i < 10; i++ {
				output1, err1 := e.Retrieve("")
				So(err1, ShouldEqual, nil)
				So(output1, ShouldEqual, info1)
			}

			patches.Reset() // if not reset will occur:patch has been existed
			info2 := "hello golang"
			patches.ApplyMethodReturn(e, "Retrieve", info2, nil)
			for i := 0; i < 10; i++ {
				output2, err2 := e.Retrieve("")
				So(err2, ShouldEqual, nil)
				So(output2, ShouldEqual, info2)
			}
		})
	})
}
