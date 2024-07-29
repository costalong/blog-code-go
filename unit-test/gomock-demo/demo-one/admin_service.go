package demo_one

import "strconv"

//go:generate  mockgen -destination=mock/admin_service_mock.go  -package=mock  -aux_files=github.com/costalong/blog-code-go/unit-test/gomock-demo/demo-one=admin_info.go -source=admin_service.go
type AdminSrv interface {
	GetUserById(id int) string
	Update(id int, data interface{}) (string, error)
	AdminInfoSrv
}

type AdminService struct {
}

func NewAdminService() *UserService {
	return &UserService{}
}

func (s *AdminService) GetUserById(id int) string {
	return strconv.Itoa(id)
}

func (s *AdminService) Update(id int, data interface{}) (string, error) {
	str := s.GetUserById(id)
	return str, nil
}
